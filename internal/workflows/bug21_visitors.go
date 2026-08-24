package workflows

var visitorArrivalCounts = map[string]int{}

func CountVisitorArrivals(visitors []string) map[string]int {
	for _, visitor := range visitors {
		visitorArrivalCounts[visitor]++
	}
	return visitorArrivalCounts
}
