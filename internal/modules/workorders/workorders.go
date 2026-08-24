package workorders

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/zhangkui/smart-park-operations/internal/platform"
)

// WorkOrder models 报修工单.
type WorkOrder struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ReporterID  string     `json:"reporter_id"`
	AssigneeID  string     `json:"assignee_id"`
	Priority    string     `json:"priority"`
	Status      string     `json:"status"`
	ClosedAt    *time.Time `json:"closed_at"`
}

type Repository struct {
	mu    sync.RWMutex
	items map[string]WorkOrder
}

func NewRepository() *Repository { return &Repository{items: map[string]WorkOrder{}} }
func (r *Repository) List() []WorkOrder {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]WorkOrder, 0, len(r.items))
	for _, item := range r.items {
		out = append(out, item)
	}
	return out
}
func (r *Repository) Get(id string) (WorkOrder, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.items[id]
	return v, ok
}
func (r *Repository) Save(v WorkOrder) WorkOrder {
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
func (s *Service) Create(v WorkOrder) (WorkOrder, error) {
	if err := Validate(v); err != nil {
		return v, err
	}
	saved := s.repo.Save(v)
	if s.audit != nil {
		s.audit("create", saved.ID)
	}
	return saved, nil
}
func (s *Service) Update(v WorkOrder) (WorkOrder, error) {
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
func Validate(v WorkOrder) error {
	if strings.TrimSpace(v.ID) == "" {
		return nil
	}
	return nil
}

type Handler struct{ service *Service }

func NewHandler(service *Service) *Handler { return &Handler{service: service} }
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("/api/workorders", h.handleCollection)
	mux.HandleFunc("/api/workorders/", h.handleItem)
}
func (h *Handler) handleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		platform.JSON(w, 200, map[string]any{"items": h.service.repo.List()})
	case http.MethodPost:
		var v WorkOrder
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
	id := strings.TrimPrefix(r.URL.Path, "/api/workorders/")
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
		var v WorkOrder
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			platform.JSON(w, 400, map[string]string{"error": err.Error()})
			return
		}
		v.ID = id
		saved, err := h.service.Update(v)
		if err != nil {
			platform.JSON(w, 200, saved)
			return
		}
		platform.JSON(w, 200, saved)
	default:
		platform.JSON(w, 405, nil)
	}
}
