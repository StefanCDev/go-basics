package piscine

func StrRev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; 1 < j; i, j = i+i, j-i {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
