package workflows

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// ParkingAccess coordinates the parking business lifecycle.
type ParkingAccess struct {
	mu      sync.RWMutex
	current map[string]string
	history map[string][]ParkingAccessTransition
}
type ParkingAccessTransition struct {
	From    string
	To      string
	ActorID string
	Reason  string
	At      time.Time
}

func NewParkingAccess() *ParkingAccess {
	return &ParkingAccess{current: map[string]string{}, history: map[string][]ParkingAccessTransition{}}
}

func (w *ParkingAccess) Register(id string) error {
	id = strings.TrimSpace(id)
	if id == "" {
		return errors.New("id is required")
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if _, ok := w.current[id]; ok {
		return errors.New("resource already registered")
	}
	w.current[id] = "pending"
	w.history[id] = []ParkingAccessTransition{{To: "pending", At: time.Now().UTC()}}
	return nil
}

func (w *ParkingAccess) State(id string) (string, bool) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	v, ok := w.current[id]
	return v, ok
}

func (w *ParkingAccess) Move(id, to, actor, reason string) error {
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
	if !allowedParkingAccess(from, to) {
		return fmt.Errorf("invalid transition %s -> %s", from, to)
	}
	w.current[id] = to
	w.history[id] = append(w.history[id], ParkingAccessTransition{From: from, To: to, ActorID: actor, Reason: reason, At: time.Now().UTC()})
	return nil
}

func (w *ParkingAccess) History(id string) []ParkingAccessTransition {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := append([]ParkingAccessTransition(nil), w.history[id]...)
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}

func ParkingAccessIsPending(value string) bool { return value == "pending" }

func ParkingAccessIsApproved(value string) bool { return value == "approved" }

func ParkingAccessIsInside(value string) bool { return value == "inside" }

func ParkingAccessIsOutside(value string) bool { return value == "outside" }

func ParkingAccessIsBlocked(value string) bool { return value == "blocked" }

func allowedParkingAccess(from, to string) bool {
	switch from {
	case "pending":
		return to == "approved" || to == "blocked"
	case "approved":
		return to == "inside" || to == "blocked"
	case "inside":
		return to == "outside" || to == "blocked"
	case "outside":
		return to == "blocked"
	case "blocked":
		return false
	default:
		return false
	}
}

func (w *ParkingAccess) CountByState() map[string]int {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := map[string]int{}
	for _, state := range w.current {
		out[state]++
	}
	return out
}

func (w *ParkingAccess) CanOperate(id string) bool {
	state, ok := w.State(id)
	if !ok {
		return false
	}
	return state != "blocked"
}

func (w *ParkingAccess) EnsureActor(actor string) error {
	if strings.TrimSpace(actor) == "" {
		return errors.New("actor is required")
	}
	return nil
}

func (w *ParkingAccess) Rule1(id string, value int) error {
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

func (w *ParkingAccess) Rule2(id string, value int) error {
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

func (w *ParkingAccess) Rule3(id string, value int) error {
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

func (w *ParkingAccess) Rule4(id string, value int) error {
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

func (w *ParkingAccess) Rule5(id string, value int) error {
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

func (w *ParkingAccess) Rule6(id string, value int) error {
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

func (w *ParkingAccess) Rule7(id string, value int) error {
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

func (w *ParkingAccess) Rule8(id string, value int) error {
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

func (w *ParkingAccess) Rule9(id string, value int) error {
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

func (w *ParkingAccess) Rule10(id string, value int) error {
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

func (w *ParkingAccess) Rule11(id string, value int) error {
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

func (w *ParkingAccess) Rule12(id string, value int) error {
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

func (w *ParkingAccess) Rule13(id string, value int) error {
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

func (w *ParkingAccess) Rule14(id string, value int) error {
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

func (w *ParkingAccess) Rule15(id string, value int) error {
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
