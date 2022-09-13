package piscine

func Sqrt(nb int) int {
	answer := 0
	if nb == 1 {
		return 1
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 && i*i == nb {
			answer = i
		}
	}
	return answer
}
