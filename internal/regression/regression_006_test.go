package regression

import (
	"testing"

	"github.com/zhangkui/smart-park-operations/internal/modules/audit"
)

func TestBug06_AuditLogRejectsMissingRetentionFields(t *testing.T) {
	_, err := audit.NewService(audit.NewRepository(), nil).Create(audit.AuditLog{Action: "delete"})
	if err == nil {
		t.Fatal("expected invalid audit log to be rejected")
	}
}
