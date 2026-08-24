package workflows

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// TenantLease coordinates the tenants business lifecycle.
type TenantLease struct {
	mu      sync.RWMutex
	current map[string]string
	history map[string][]TenantLeaseTransition
}
type TenantLeaseTransition struct {
	From    string
	To      string
	ActorID string
	Reason  string
	At      time.Time
}

func NewTenantLease() *TenantLease {
	return &TenantLease{current: map[string]string{}, history: map[string][]TenantLeaseTransition{}}
}

func (w *TenantLease) Register(id string) error {
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
	w.history[id] = []TenantLeaseTransition{{To: "pending", At: time.Now().UTC()}}
	return nil
}

func (w *TenantLease) State(id string) (string, bool) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	v, ok := w.current[id]
	return v, ok
}

func (w *TenantLease) Move(id, to, actor, reason string) error {
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
	if !allowedTenantLease(from, to) {
		return fmt.Errorf("invalid transition %s -> %s", from, to)
	}
	w.current[id] = to
	w.history[id] = append(w.history[id], TenantLeaseTransition{From: from, To: to, ActorID: actor, Reason: reason, At: time.Now().UTC()})
	return nil
}

func (w *TenantLease) History(id string) []TenantLeaseTransition {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := append([]TenantLeaseTransition(nil), w.history[id]...)
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}

func TenantLeaseIsPending(value string) bool { return value == "pending" }

func TenantLeaseIsActive(value string) bool { return value == "active" }

func TenantLeaseIsSuspended(value string) bool { return value == "suspended" }

func TenantLeaseIsExpired(value string) bool { return value == "expired" }

func allowedTenantLease(from, to string) bool {
	switch from {
	case "pending":
		return to == "active" || to == "expired"
	case "active":
		return to == "suspended" || to == "expired"
	case "suspended":
		return to == "expired"
	case "expired":
		return false
	default:
		return false
	}
}

func (w *TenantLease) CountByState() map[string]int {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := map[string]int{}
	for _, state := range w.current {
		out[state]++
	}
	return out
}

func (w *TenantLease) CanOperate(id string) bool {
	state, ok := w.State(id)
	if !ok {
		return false
	}
	return state != "expired"
}

func (w *TenantLease) EnsureActor(actor string) error {
	if strings.TrimSpace(actor) == "" {
		return errors.New("actor is required")
	}
	return nil
}

func (w *TenantLease) Rule1(id string, value int) error {
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

func (w *TenantLease) Rule2(id string, value int) error {
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

func (w *TenantLease) Rule3(id string, value int) error {
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

func (w *TenantLease) Rule4(id string, value int) error {
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

func (w *TenantLease) Rule5(id string, value int) error {
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

func (w *TenantLease) Rule6(id string, value int) error {
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

func (w *TenantLease) Rule7(id string, value int) error {
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

func (w *TenantLease) Rule8(id string, value int) error {
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

func (w *TenantLease) Rule9(id string, value int) error {
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

func (w *TenantLease) Rule10(id string, value int) error {
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

func (w *TenantLease) Rule11(id string, value int) error {
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

func (w *TenantLease) Rule12(id string, value int) error {
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

func (w *TenantLease) Rule13(id string, value int) error {
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

func (w *TenantLease) Rule14(id string, value int) error {
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

func (w *TenantLease) Rule15(id string, value int) error {
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
