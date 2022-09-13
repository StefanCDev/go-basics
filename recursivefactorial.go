package piscine

func RecursiveFactorial(nb int) int {
	if nb > 21 {
		return 0
	}
	if nb < 1 {
		return 0
	}
	if nb <= 1 {
		return 1
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}
