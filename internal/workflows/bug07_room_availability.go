package workflows

import "runtime"

// RoomAvailability tracks whether a meeting room has been allocated.
type RoomAvailability struct {
	reserved map[string]bool
}

func NewRoomAvailability() *RoomAvailability {
	return &RoomAvailability{reserved: map[string]bool{}}
}

func (availability *RoomAvailability) TryReserve(roomID string) bool {
	if availability.reserved[roomID] {
		return false
	}
	runtime.Gosched()
	availability.reserved[roomID] = true
	return true
}
