package regression

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zhangkui/smart-park-operations/internal/modules/workorders"
)

func TestBug04_WorkOrderUpdateReturnsRepositoryErrors(t *testing.T) {
	mux := http.NewServeMux()
	workorders.NewHandler(workorders.NewService(workorders.NewRepository(), nil)).Routes(mux)
	request := httptest.NewRequest(http.MethodPut, "/api/workorders/missing", bytes.NewBufferString(`{"title":"broken"}`))
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected repository error status 422, got %d", recorder.Code)
	}
}
