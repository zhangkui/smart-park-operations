package regression

import (
	"context"
	"testing"

	"github.com/zhangkui/smart-park-operations/internal/workflows"
)

func TestBug10_CancelledVisitorVerificationSkipsAuditWrite(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	writes := 0
	err := workflows.VerifyVisitorWithAudit(ctx, "visitor-01", func(string) { writes++ })
	if err == nil || writes != 0 {
		t.Fatalf("expected no audit write after cancellation, got writes=%d err=%v", writes, err)
	}
}
