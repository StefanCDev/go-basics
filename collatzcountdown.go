package piscine

func CollatzCountdown(start int) int {
	var stepnb int
	if start <= 0 {
		return -1
	}
	for start > 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = (3 * start) + 1
		}
		stepnb++
	}
	return stepnb
}
