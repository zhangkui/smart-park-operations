package workflows

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// EnergyAggregation coordinates the energy business lifecycle.
type EnergyAggregation struct {
	mu      sync.RWMutex
	current map[string]string
	history map[string][]EnergyAggregationTransition
}
type EnergyAggregationTransition struct {
	From    string
	To      string
	ActorID string
	Reason  string
	At      time.Time
}

func NewEnergyAggregation() *EnergyAggregation {
	return &EnergyAggregation{current: map[string]string{}, history: map[string][]EnergyAggregationTransition{}}
}

func (w *EnergyAggregation) Register(id string) error {
	id = strings.TrimSpace(id)
	if id == "" {
		return errors.New("id is required")
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if _, ok := w.current[id]; ok {
		return errors.New("resource already registered")
	}
	w.current[id] = "raw"
	w.history[id] = []EnergyAggregationTransition{{To: "raw", At: time.Now().UTC()}}
	return nil
}

func (w *EnergyAggregation) State(id string) (string, bool) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	v, ok := w.current[id]
	return v, ok
}

func (w *EnergyAggregation) Move(id, to, actor, reason string) error {
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
	if !allowedEnergyAggregation(from, to) {
		return fmt.Errorf("invalid transition %s -> %s", from, to)
	}
	w.current[id] = to
	w.history[id] = append(w.history[id], EnergyAggregationTransition{From: from, To: to, ActorID: actor, Reason: reason, At: time.Now().UTC()})
	return nil
}

func (w *EnergyAggregation) History(id string) []EnergyAggregationTransition {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := append([]EnergyAggregationTransition(nil), w.history[id]...)
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}

func EnergyAggregationIsRaw(value string) bool { return value == "raw" }

func EnergyAggregationIsValidated(value string) bool { return value == "validated" }

func EnergyAggregationIsAggregated(value string) bool { return value == "aggregated" }

func EnergyAggregationIsPublished(value string) bool { return value == "published" }

func allowedEnergyAggregation(from, to string) bool {
	switch from {
	case "raw":
		return to == "validated" || to == "published"
	case "validated":
		return to == "aggregated" || to == "published"
	case "aggregated":
		return to == "published"
	case "published":
		return false
	default:
		return false
	}
}

func (w *EnergyAggregation) CountByState() map[string]int {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := map[string]int{}
	for _, state := range w.current {
		out[state]++
	}
	return out
}

func (w *EnergyAggregation) CanOperate(id string) bool {
	state, ok := w.State(id)
	if !ok {
		return false
	}
	return state != "published"
}

func (w *EnergyAggregation) EnsureActor(actor string) error {
	if strings.TrimSpace(actor) == "" {
		return errors.New("actor is required")
	}
	return nil
}

func (w *EnergyAggregation) Rule1(id string, value int) error {
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

func (w *EnergyAggregation) Rule2(id string, value int) error {
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

func (w *EnergyAggregation) Rule3(id string, value int) error {
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

func (w *EnergyAggregation) Rule4(id string, value int) error {
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

func (w *EnergyAggregation) Rule5(id string, value int) error {
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

func (w *EnergyAggregation) Rule6(id string, value int) error {
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

func (w *EnergyAggregation) Rule7(id string, value int) error {
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

func (w *EnergyAggregation) Rule8(id string, value int) error {
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

func (w *EnergyAggregation) Rule9(id string, value int) error {
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

func (w *EnergyAggregation) Rule10(id string, value int) error {
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

func (w *EnergyAggregation) Rule11(id string, value int) error {
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

func (w *EnergyAggregation) Rule12(id string, value int) error {
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

func (w *EnergyAggregation) Rule13(id string, value int) error {
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

func (w *EnergyAggregation) Rule14(id string, value int) error {
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

func (w *EnergyAggregation) Rule15(id string, value int) error {
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
