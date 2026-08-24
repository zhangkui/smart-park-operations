package regression

import (
	"testing"
	"time"

	"github.com/zhangkui/smart-park-operations/internal/modules/spaces"
)

func TestBug03_SpaceBookingListDoesNotAliasRepositoryCache(t *testing.T) {
	repository := spaces.NewRepository()
	repository.Save(spaces.SpaceBooking{ID: "booking-01", SpaceID: "room-01", Status: "requested", StartAt: time.Now().UTC(), EndAt: time.Now().UTC().Add(time.Hour)})

	firstSnapshot := repository.List()
	firstSnapshot[0].Status = "cancelled"
	freshSnapshot := repository.List()
	if freshSnapshot[0].Status != "requested" {
		t.Fatalf("expected a stable booking snapshot, got mutated status %q", freshSnapshot[0].Status)
	}
}
