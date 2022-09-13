package piscine

func IterativeFactorial(nb int) int {
	if nb > 200000 || nb < 0 {
		return 0
	}
	if nb <= 1 {
		return 1
	}
	return nb * IterativeFactorial(nb-1)
}
