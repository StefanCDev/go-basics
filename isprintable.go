package piscine

func IsPrintable(s string) bool {
	sRune := []byte(s)
	for _, lens := range sRune {
		if !(lens >= 32 && lens <= 127) {
			return false
		}
	}
	return true
}
