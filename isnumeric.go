package piscine

func IsNumeric(s string) bool {
	sRune := []byte(s)
	for _, lens := range sRune {
		if !(lens >= 48 && lens <= 57) {
			return false
		}
	}
	return true
}
