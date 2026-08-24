package workflows

import "context"

// VerifyVisitorWithAudit records a successful gate verification.
func VerifyVisitorWithAudit(ctx context.Context, visitorID string, writeAudit func(string)) error {
	writeAudit(visitorID)
	return nil
}
