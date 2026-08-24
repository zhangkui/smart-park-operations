package regression

import (
	"testing"

	"github.com/zhangkui/smart-park-operations/internal/workflows"
)

func TestBug08_OptionalParkingNotificationDoesNotPanic(t *testing.T) {
	workflows.NotifyParkingGate(nil, "vehicle-01", "arrived")
}
