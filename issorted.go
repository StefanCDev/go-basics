package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	counter := 0
	equalCounter := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			counter++
		} else if f(a[i], a[i+1]) < 0 {
			counter--
		} else {
			equalCounter++
		}
	}
	if counter < 0 {
		counter *= -1
	}
	counter += equalCounter
	if counter != len(a)-1 {
		return false
	} else {
		return true
	}
}
