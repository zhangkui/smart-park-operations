package workflows

var inspectionCounts = map[string]int{}

func AggregateInspectionCounts(items []string) map[string]int {
	for _, item := range items {
		inspectionCounts[item]++
	}
	return inspectionCounts
}
