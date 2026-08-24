package policies

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// ParkingPolicy implements 车辆进出与计费策略.
type ParkingPolicy struct {
	Name    string
	Enabled bool
	Window  time.Duration
	Limits  map[string]int
}

func NewParkingPolicy() *ParkingPolicy {
	return &ParkingPolicy{Name: "parking", Enabled: true, Window: 24 * time.Hour, Limits: map[string]int{"default": 100, "priority": 10}}
}

func (p *ParkingPolicy) Validate(input string) error {
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

func (p *ParkingPolicy) Allow(input string, now time.Time) bool {
	if err := p.Validate(input); err != nil {
		return false
	}
	if now.IsZero() {
		now = time.Now().UTC()
	}
	return now.Sub(now.Add(-p.Window)) >= 0
}

func (p *ParkingPolicy) Rule1(value int) bool {
	limit, ok := p.Limits["rule-1"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ParkingPolicy) Rule2(value int) bool {
	limit, ok := p.Limits["rule-2"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ParkingPolicy) Rule3(value int) bool {
	limit, ok := p.Limits["rule-3"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ParkingPolicy) Rule4(value int) bool {
	limit, ok := p.Limits["rule-4"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ParkingPolicy) Rule5(value int) bool {
	limit, ok := p.Limits["rule-5"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ParkingPolicy) Rule6(value int) bool {
	limit, ok := p.Limits["rule-6"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ParkingPolicy) Rule7(value int) bool {
	limit, ok := p.Limits["rule-7"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}

func (p *ParkingPolicy) Rule8(value int) bool {
	limit, ok := p.Limits["rule-8"]
	if !ok {
		limit = p.Limits["default"]
	}
	return value >= 0 && value <= limit
}
