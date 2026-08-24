package policies

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// ReservationPolicy implements 空间预约冲突策略.
type ReservationPolicy struct {
	Name    string
	Enabled bool
	Window  time.Duration
	Limits  map[string]int
}

func NewReservationPolicy() *ReservationPolicy {
	return &ReservationPolicy{Name: "reservation", Enabled: true, Window: 24 * time.Hour, Limits: map[string]int{"default": 100, "priority": 10}}
}

func (p *ReservationPolicy) Validate(input string) error {
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

func (p *ReservationPolicy) Allow(input string, now time.Time) bool {
	if err := p.Validate(input); err != nil {
		return false
	}
	if now.IsZero() {
		now = time.Now().UTC()
	}
	return now.Sub(now.Add(-p.Window)) >= 0
}

func (p *ReservationPolicy) Rule1(value int) bool {
	limit, ok := p.Limits["rule-1"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ReservationPolicy) Rule2(value int) bool {
	limit, ok := p.Limits["rule-2"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ReservationPolicy) Rule3(value int) bool {
	limit, ok := p.Limits["rule-3"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ReservationPolicy) Rule4(value int) bool {
	limit, ok := p.Limits["rule-4"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ReservationPolicy) Rule5(value int) bool {
	limit, ok := p.Limits["rule-5"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ReservationPolicy) Rule6(value int) bool {
	limit, ok := p.Limits["rule-6"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ReservationPolicy) Rule7(value int) bool {
	limit, ok := p.Limits["rule-7"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ReservationPolicy) Rule8(value int) bool {
	limit, ok := p.Limits["rule-8"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}
