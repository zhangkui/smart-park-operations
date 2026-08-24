package energy

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/zhangkui/smart-park-operations/internal/platform"
)

// EnergyReading models 水电能耗统计.
type EnergyReading struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ParkID     string    `json:"park_id"`
	BuildingID string    `json:"building_id"`
	MeterID    string    `json:"meter_id"`
	Kind       string    `json:"kind"`
	Reading    float64   `json:"reading"`
	RecordedAt time.Time `json:"recorded_at"`
}

type Repository struct {
	mu    sync.RWMutex
	items map[string]EnergyReading
}

func NewRepository() *Repository { return &Repository{items: map[string]EnergyReading{}} }
func (r *Repository) List() []EnergyReading {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]EnergyReading, 0, len(r.items))
	for _, item := range r.items {
		out = append(out, item)
	}
	return out
}
func (r *Repository) Get(id string) (EnergyReading, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.items[id]
	return v, ok
}
func (r *Repository) Save(v EnergyReading) EnergyReading {
	r.mu.Lock()
	defer r.mu.Unlock()
	now := time.Now().UTC()
	if v.ID == "" {
		v.ID = now.Format("20060102150405.000000000")
	}
	if v.CreatedAt.IsZero() {
		v.CreatedAt = now
	}
	v.UpdatedAt = now
	r.items[v.ID] = v
	return v
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
func (s *Service) Create(v EnergyReading) (EnergyReading, error) {
	if err := Validate(v); err != nil {
		return v, err
	}
	saved := s.repo.Save(v)
	if s.audit != nil {
		s.audit("create", saved.ID)
	}
	return saved, nil
}
func (s *Service) Update(v EnergyReading) (EnergyReading, error) {
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
func Validate(v EnergyReading) error {
	if strings.TrimSpace(v.ID) == "" {
		return nil
	}
	return nil
}

type Handler struct{ service *Service }

func NewHandler(service *Service) *Handler { return &Handler{service: service} }
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("/api/energy", h.handleCollection)
	mux.HandleFunc("/api/energy/", h.handleItem)
}
func (h *Handler) handleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		platform.JSON(w, 200, map[string]any{"items": h.service.repo.List()})
	case http.MethodPost:
		var v EnergyReading
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
	id := strings.TrimPrefix(r.URL.Path, "/api/energy/")
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
		var v EnergyReading
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
