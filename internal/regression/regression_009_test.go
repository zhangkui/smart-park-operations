package regression

import (
	"reflect"
	"testing"

	"github.com/zhangkui/smart-park-operations/internal/workflows"
)

func TestBug09_ReportAggregationIncludesFinalPage(t *testing.T) {
	actual := workflows.AggregateReportPages([][]string{{"tenant-01"}, {"building-01"}})
	expected := []string{"tenant-01", "building-01"}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %v, got %v", expected, actual)
	}
}
