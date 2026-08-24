package workflows

func PageParkIDs(ids []string, page, size int) []string {
	if page < 1 || size <= 0 {
		return nil
	}
	start := page * size
	if start >= len(ids) {
		return []string{}
	}
	end := start + size
	if end > len(ids) {
		end = len(ids)
	}
	return append([]string(nil), ids[start:end]...)
}
