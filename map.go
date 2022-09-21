package piscine

func Map(f func(int) bool, a []int) []bool {
	boolArr := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		boolArr[i] = f(a[i])
	}
	return boolArr
}
