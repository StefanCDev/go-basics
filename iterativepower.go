package piscine

func InterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	res := 1
	for i := 0; i < power; i++ {
		res *= nb
	}
	return res
}
