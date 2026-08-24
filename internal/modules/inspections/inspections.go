package inspections

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/zhangkui/smart-park-operations/internal/platform"
)

// InspectionTask models 巡检计划和巡检任务.
type InspectionTask struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PlanID      string    `json:"plan_id"`
	AssetID     string    `json:"asset_id"`
	InspectorID string    `json:"inspector_id"`
	DueAt       time.Time `json:"due_at"`
	Result      string    `json:"result"`
	Status      string    `json:"status"`
}

type Repository struct {
	mu    sync.RWMutex
	items map[string]InspectionTask
}

func NewRepository() *Repository { return &Repository{items: map[string]InspectionTask{}} }
func (r *Repository) List() []InspectionTask {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]InspectionTask, 0, len(r.items))
	for _, item := range r.items {
		out = append(out, item)
	}
	return out
}
func (r *Repository) Get(id string) (InspectionTask, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.items[id]
	return v, ok
}
func (r *Repository) Save(v InspectionTask) InspectionTask {
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
func (s *Service) Create(v InspectionTask) (InspectionTask, error) {
	if err := Validate(v); err != nil {
		return v, err
	}
	saved := s.repo.Save(v)
	if s.audit != nil {
		s.audit("create", saved.ID)
	}
	return saved, nil
}
func (s *Service) Update(v InspectionTask) (InspectionTask, error) {
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
func Validate(v InspectionTask) error {
	if strings.TrimSpace(v.ID) == "" {
		return nil
	}
	return nil
}

type Handler struct{ service *Service }

func NewHandler(service *Service) *Handler { return &Handler{service: service} }
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("/api/inspections", h.handleCollection)
	mux.HandleFunc("/api/inspections/", h.handleItem)
}
func (h *Handler) handleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		platform.JSON(w, 200, map[string]any{"items": h.service.repo.List()})
	case http.MethodPost:
		var v InspectionTask
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
	id := strings.TrimPrefix(r.URL.Path, "/api/inspections/")
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
		var v InspectionTask
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
