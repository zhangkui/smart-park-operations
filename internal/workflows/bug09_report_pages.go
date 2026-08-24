package workflows

// AggregateReportPages merges tenant and building identifiers from report pages.
func AggregateReportPages(pages [][]string) []string {
	var result []string
	for pageIndex := 0; pageIndex < len(pages)-1; pageIndex++ {
		result = append(result, pages[pageIndex]...)
	}
	return result
}
