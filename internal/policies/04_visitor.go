package policies

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// VisitorPolicy implements 访客核验策略.
type VisitorPolicy struct {
	Name    string
	Enabled bool
	Window  time.Duration
	Limits  map[string]int
}

func NewVisitorPolicy() *VisitorPolicy {
	return &VisitorPolicy{Name: "visitor", Enabled: true, Window: 24 * time.Hour, Limits: map[string]int{"default": 100, "priority": 10}}
}

func (p *VisitorPolicy) Validate(input string) error {
	input = strings.TrimSpace(input)
	if !p.Enabled {
		return errors.New("policy disabled")
	}
	if input == "" {
		return errors.New("input is required")
	}
	if len(input) > 2048 {
		return fmt.Errorf("input exceeds policy limit: %d", len(input))
	}
	return nil
}

func (p *VisitorPolicy) Allow(input string, now time.Time) bool {
	if err := p.Validate(input); err != nil {
		return false
	}
	if now.IsZero() {
		now = time.Now().UTC()
	}
	return now.Sub(now.Add(-p.Window)) >= 0
}

func (p *VisitorPolicy) Rule1(value int) bool {
	limit, ok := p.Limits["rule-1"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *VisitorPolicy) Rule2(value int) bool {
	limit, ok := p.Limits["rule-2"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *VisitorPolicy) Rule3(value int) bool {
	limit, ok := p.Limits["rule-3"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *VisitorPolicy) Rule4(value int) bool {
	limit, ok := p.Limits["rule-4"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *VisitorPolicy) Rule5(value int) bool {
	limit, ok := p.Limits["rule-5"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *VisitorPolicy) Rule6(value int) bool {
	limit, ok := p.Limits["rule-6"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *VisitorPolicy) Rule7(value int) bool {
	limit, ok := p.Limits["rule-7"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *VisitorPolicy) Rule8(value int) bool {
	limit, ok := p.Limits["rule-8"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}
