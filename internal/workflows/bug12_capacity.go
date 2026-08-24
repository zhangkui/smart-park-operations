package workflows

func CanAllocateCapacity(current, requested, capacity int) bool {
	if current < 0 || requested <= 0 || capacity < 0 {
		return false
	}
	return current+requested < capacity
}
