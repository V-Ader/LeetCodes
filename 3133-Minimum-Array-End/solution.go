func minEnd(size int, startValue int) int64 {
	size--
	for i := 1; size > 0; i <<= 1 {
		if startValue&i != i {
			if size&1 == 1 {
				startValue = startValue | i
			}
			size = size >> 1
		}
	}
	return int64(startValue)
}