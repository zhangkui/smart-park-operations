package workflows

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// ReportLifecycle coordinates the reports business lifecycle.
type ReportLifecycle struct {
	mu      sync.RWMutex
	current map[string]string
	history map[string][]ReportLifecycleTransition
}
type ReportLifecycleTransition struct {
	From    string
	To      string
	ActorID string
	Reason  string
	At      time.Time
}

func NewReportLifecycle() *ReportLifecycle {
	return &ReportLifecycle{current: map[string]string{}, history: map[string][]ReportLifecycleTransition{}}
}

func (w *ReportLifecycle) Register(id string) error {
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
	w.history[id] = []ReportLifecycleTransition{{To: "draft", At: time.Now().UTC()}}
	return nil
}

func (w *ReportLifecycle) State(id string) (string, bool) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	v, ok := w.current[id]
	return v, ok
}

func (w *ReportLifecycle) Move(id, to, actor, reason string) error {
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
	if !allowedReportLifecycle(from, to) {
		return fmt.Errorf("invalid transition %s -> %s", from, to)
	}
	w.current[id] = to
	w.history[id] = append(w.history[id], ReportLifecycleTransition{From: from, To: to, ActorID: actor, Reason: reason, At: time.Now().UTC()})
	return nil
}

func (w *ReportLifecycle) History(id string) []ReportLifecycleTransition {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := append([]ReportLifecycleTransition(nil), w.history[id]...)
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}

func ReportLifecycleIsDraft(value string) bool { return value == "draft" }

func ReportLifecycleIsGenerating(value string) bool { return value == "generating" }

func ReportLifecycleIsPublished(value string) bool { return value == "published" }

func ReportLifecycleIsExpired(value string) bool { return value == "expired" }

func allowedReportLifecycle(from, to string) bool {
	switch from {
	case "draft":
		return to == "generating" || to == "expired"
	case "generating":
		return to == "published" || to == "expired"
	case "published":
		return to == "expired"
	case "expired":
		return false
	default:
		return false
	}
}

func (w *ReportLifecycle) CountByState() map[string]int {
	w.mu.RLock()
	defer w.mu.RUnlock()
	out := map[string]int{}
	for _, state := range w.current {
		out[state]++
	}
	return out
}

func (w *ReportLifecycle) CanOperate(id string) bool {
	state, ok := w.State(id)
	if !ok {
		return false
	}
	return state != "expired"
}

func (w *ReportLifecycle) EnsureActor(actor string) error {
	if strings.TrimSpace(actor) == "" {
		return errors.New("actor is required")
	}
	return nil
}

func (w *ReportLifecycle) Rule1(id string, value int) error {
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

func (w *ReportLifecycle) Rule2(id string, value int) error {
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

func (w *ReportLifecycle) Rule3(id string, value int) error {
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

func (w *ReportLifecycle) Rule4(id string, value int) error {
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

func (w *ReportLifecycle) Rule5(id string, value int) error {
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

func (w *ReportLifecycle) Rule6(id string, value int) error {
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

func (w *ReportLifecycle) Rule7(id string, value int) error {
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

func (w *ReportLifecycle) Rule8(id string, value int) error {
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

func (w *ReportLifecycle) Rule9(id string, value int) error {
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

func (w *ReportLifecycle) Rule10(id string, value int) error {
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

func (w *ReportLifecycle) Rule11(id string, value int) error {
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

func (w *ReportLifecycle) Rule12(id string, value int) error {
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

func (w *ReportLifecycle) Rule13(id string, value int) error {
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

func (w *ReportLifecycle) Rule14(id string, value int) error {
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

func (w *ReportLifecycle) Rule15(id string, value int) error {
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
