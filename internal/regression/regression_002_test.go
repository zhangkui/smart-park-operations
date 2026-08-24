package regression

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/zhangkui/smart-park-operations/internal/modules/visitors"
)

func TestBug02_VisitorCreateAllowsOptionalAuditCallback(t *testing.T) {
	service := visitors.NewService(visitors.NewRepository(), nil)
	handler := visitors.NewHandler(service)
	mux := http.NewServeMux()
	handler.Routes(mux)

	payload := []byte(`{"name":"Lin","phone":"13800138000","tenant_id":"tenant-01","host_id":"host-01","visit_at":"` + time.Now().UTC().Format(time.RFC3339) + `"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/visitors", bytes.NewReader(payload))
	recorder := httptest.NewRecorder()

	mux.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected visitor appointment creation to succeed without an audit callback, got status %d: %s", recorder.Code, recorder.Body.String())
	}
}
