package regression

import (
	"fmt"
	"sync"
	"testing"

	"github.com/zhangkui/smart-park-operations/internal/modules/parking"
	"github.com/zhangkui/smart-park-operations/internal/workflows"
)

func TestBug01_GateEventCacheConcurrent(t *testing.T) {
	repository := parking.NewRepository()
	service := parking.NewService(repository, nil)
	lifecycle := workflows.NewParkingAccess()
	service.BindAccessLifecycle(lifecycle)

	const events = 32
	for index := 0; index < events; index++ {
		id := fmt.Sprintf("vehicle-%d", index)
		repository.Save(parking.Vehicle{ID: id, PlateNumber: id, AccessStatus: "approved"})
		if err := lifecycle.Register(id); err != nil {
			t.Fatal(err)
		}
		if err := lifecycle.Move(id, "approved", "system", "approval"); err != nil {
			t.Fatal(err)
		}
	}

	var group sync.WaitGroup
	for index := 0; index < events; index++ {
		id := fmt.Sprintf("vehicle-%d", index)
		group.Add(1)
		go func() {
			defer group.Done()
			if _, err := service.RecordGateEvent(id, "inside", "gate-1"); err != nil {
				t.Errorf("record gate event: %v", err)
			}
		}()
	}
	group.Wait()
}
