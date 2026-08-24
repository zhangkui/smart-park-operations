package workflows

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// ParkOccupancy coordinates the parks business lifecycle.
type ParkOccupancy struct {
	mu      sync.RWMutex
	current map[string]string
	history map[string][]ParkOccupancyTransition
}
type ParkOccupancyTransition struct {
	From    string
	To      string
	ActorID string
	Reason  string
	At      time.Time
}

func NewParkOccupancy() *ParkOccupancy {
	return &ParkOccupancy{current: map[string]string{}, history: map[string][]ParkOccupancyTransition{}}
}

func (w *ParkOccupancy) Register(id string) error {
	id = strings.TrimSpace(id)
	if id == "" {
		return errors.New("id is required")
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if _, ok := w.current[id]; ok {
		return errors.New("resource already registered")
	}
	w.current[id] = "draft"
	w.history[id] = []ParkOccupancyTransition{{To: "draft", At: time.Now().UTC()}}
	return nil
}

func (w *ParkOccupancy) State(id string) (string, bool) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	v, ok := w.current[id]
	return v, ok
}

func (w *ParkOccupancy) Move(id, to, actor, reason string) error {
	to = strings.TrimSpace(to)
	if to == "" {
		return errors.New("target state is required")
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	from, ok := w.current[id]
	if !ok {
		return errors.New("resource is not registered")
	}
	if !allowedParkOccupancy(from, to) {
		return fmt.Errorf("invalid transition %s -> %s", from, to)
	}
	w.current[id] = to
	w.history[id] = append(w.history[id], ParkOccupancyTransition{From: from, To: to, ActorID: actor, Reason: reason, At: time.Now().UTC()})
	return nil
}

func (w *ParkOccupancy) History(id string) []ParkOccupancyTransition {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := append([]ParkOccupancyTransition(nil), w.history[id]...)
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}

func ParkOccupancyIsDraft(value string) bool { return value == "draft" }

func ParkOccupancyIsActive(value string) bool { return value == "active" }

func ParkOccupancyIsMaintenance(value string) bool { return value == "maintenance" }

func ParkOccupancyIsArchived(value string) bool { return value == "archived" }

func allowedParkOccupancy(from, to string) bool {
	switch from {
	case "draft":
		return to == "active" || to == "archived"
	case "active":
		return to == "maintenance" || to == "archived"
	case "maintenance":
		return to == "archived"
	case "archived":
		return false
	default:
		return false
	}
}

func (w *ParkOccupancy) CountByState() map[string]int {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := map[string]int{}
	for _, state := range w.current {
		out[state]++
	}
	return out
}

func (w *ParkOccupancy) CanOperate(id string) bool {
	state, ok := w.State(id)
	if !ok {
		return false
	}
	return state != "archived"
}

func (w *ParkOccupancy) EnsureActor(actor string) error {
	if strings.TrimSpace(actor) == "" {
		return errors.New("actor is required")
	}
	return nil
}

func (w *ParkOccupancy) Rule1(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule2(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule3(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule4(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule5(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule6(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule7(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule8(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule9(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule10(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule11(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule12(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule13(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule14(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}

func (w *ParkOccupancy) Rule15(id string, value int) error {
	if err := w.EnsureActor(id); err != nil {
		return err
	}
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	if value > 1000000 {
		return errors.New("value exceeds operational limit")
	}
	return nil
}
