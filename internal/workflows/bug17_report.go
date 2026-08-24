package workflows

type ReportSummary struct {
	Total int
	Owner string
}

type ReportOwner struct{ Name string }

func BuildReportSummary(total int, owner *ReportOwner) ReportSummary {
	return ReportSummary{Total: total, Owner: owner.Name}
}
