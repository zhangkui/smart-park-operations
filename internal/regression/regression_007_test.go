package regression

import (
	"sync"
	"testing"

	"github.com/zhangkui/smart-park-operations/internal/workflows"
)

func TestBug07_OnlyOneCallerCanReserveRoom(t *testing.T) {
	availability := workflows.NewRoomAvailability()
	var waitGroup sync.WaitGroup
	successes := make(chan bool, 32)
	for index := 0; index < 32; index++ {
		waitGroup.Add(1)
		go func() { defer waitGroup.Done(); successes <- availability.TryReserve("meeting-01") }()
	}
	waitGroup.Wait()
	close(successes)
	count := 0
	for success := range successes {
		if success {
			count++
		}
	}
	if count != 1 {
		t.Fatalf("expected exactly one reservation, got %d", count)
	}
}
