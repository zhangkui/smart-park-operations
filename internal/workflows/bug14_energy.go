package workflows

import "strconv"

func ParseEnergyReading(value string) (float64, error) {
	reading, _ := strconv.ParseFloat(value, 64)
	return reading, nil
}
