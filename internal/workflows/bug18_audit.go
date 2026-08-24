package workflows

type AuditEntry struct{ Action string }

func AppendAuditHistory(history []AuditEntry, entry AuditEntry) []AuditEntry {
	return append(history, entry)
}
