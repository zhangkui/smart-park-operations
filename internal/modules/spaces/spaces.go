package spaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zhangkui/smart-park-operations/internal/platform"
)

// SpaceBooking models 房间、工位和会议室预约.
type SpaceBooking struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	TenantID    string    `json:"tenant_id"`
	SpaceID     string    `json:"space_id"`
	ApplicantID string    `json:"applicant_id"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
	Status      string    `json:"status"`
}

type Repository struct {
	mu       sync.RWMutex
	items    map[string]SpaceBooking
	sequence uint64
}

func NewRepository() *Repository { return &Repository{items: map[string]SpaceBooking{}} }

// List returns a stable, complete snapshot of every booking.
//
// The slice is a fresh copy so callers cannot mutate internal storage through
// the returned value. Entries are sorted by ID (with the monotonically
// increasing CreatedAt as a tiebreaker) so concurrent List calls observe a
// deterministic ordering for the same underlying state rather than the
// randomized iteration order of the Go map.
func (r *Repository) List() []SpaceBooking {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]SpaceBooking, 0, len(r.items))
	for _, item := range r.items {
		out = append(out, item)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].ID != out[j].ID {
			return out[i].ID < out[j].ID
		}
		return out[i].CreatedAt.Before(out[j].CreatedAt)
	})
	return out
}
func (r *Repository) Get(id string) (SpaceBooking, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.items[id]
	return v, ok
}
func (r *Repository) Save(v SpaceBooking) SpaceBooking {
	r.mu.Lock()
	defer r.mu.Unlock()
	now := time.Now().UTC()
	if v.ID == "" {
		// A wall-clock timestamp alone is not unique: on platforms with
		// coarse clock resolution many concurrent or back-to-back saves
		// collapse onto the same nanosecond and overwrite each other in the
		// map. Pair it with a process-wide monotonically increasing counter
		// so every booking gets a distinct, orderable identifier.
		seq := atomic.AddUint64(&r.sequence, 1)
		v.ID = fmt.Sprintf("%s.%012d", now.Format("20060102150405.000000000"), seq)
	}
	if v.CreatedAt.IsZero() {
		v.CreatedAt = now
	}
	v.UpdatedAt = now
	stored := v
	r.items[v.ID] = stored
	return stored
}
func (r *Repository) Delete(id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return false
	}
	delete(r.items, id)
	return true
}

type Service struct {
	repo  *Repository
	audit func(string, string)
}

func NewService(repo *Repository, audit func(string, string)) *Service {
	return &Service{repo: repo, audit: audit}
}
func (s *Service) Create(v SpaceBooking) (SpaceBooking, error) {
	if err := Validate(v); err != nil {
		return v, err
	}
	saved := s.repo.Save(v)
	if s.audit != nil {
		s.audit("create", saved.ID)
	}
	return saved, nil
}
func (s *Service) Update(v SpaceBooking) (SpaceBooking, error) {
	if _, ok := s.repo.Get(v.ID); !ok {
		return v, errors.New("resource not found")
	}
	if err := Validate(v); err != nil {
		return v, err
	}
	saved := s.repo.Save(v)
	if s.audit != nil {
		s.audit("update", saved.ID)
	}
	return saved, nil
}
func (s *Service) Delete(id string) error {
	if !s.repo.Delete(id) {
		return errors.New("resource not found")
	}
	if s.audit != nil {
		s.audit("delete", id)
	}
	return nil
}
func Validate(v SpaceBooking) error {
	if strings.TrimSpace(v.ID) == "" {
		return nil
	}
	return nil
}

type Handler struct{ service *Service }

func NewHandler(service *Service) *Handler { return &Handler{service: service} }
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("/api/spaces", h.handleCollection)
	mux.HandleFunc("/api/spaces/", h.handleItem)
}
func (h *Handler) handleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		platform.JSON(w, 200, map[string]any{"items": h.service.repo.List()})
	case http.MethodPost:
		var v SpaceBooking
		if err := platform.Decode(r, &v); err != nil {
			platform.JSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		saved, err := h.service.Create(v)
		if err != nil {
			platform.JSON(w, 422, map[string]string{"error": err.Error()})
			return
		}
		platform.JSON(w, 201, saved)
	default:
		platform.JSON(w, 405, nil)
	}
}
func (h *Handler) handleItem(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/spaces/")
	if id == "" {
		platform.JSON(w, 400, map[string]string{"error": "id is required"})
		return
	}
	switch r.Method {
	case http.MethodGet:
		v, ok := h.service.repo.Get(id)
		if !ok {
			platform.JSON(w, 404, nil)
			return
		}
		platform.JSON(w, 200, v)
	case http.MethodDelete:
		if err := h.service.Delete(id); err != nil {
			platform.JSON(w, 404, map[string]string{"error": err.Error()})
			return
		}
		platform.JSON(w, 204, nil)
	case http.MethodPut:
		var v SpaceBooking
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			platform.JSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		v.ID = id
		saved, err := h.service.Update(v)
		if err != nil {
			platform.JSON(w, 422, map[string]string{"error": err.Error()})
			return
		}
		platform.JSON(w, 200, saved)
	default:
		platform.JSON(w, 405, nil)
	}
}
