package piscine

func IsUpper(s string) bool {
	strSlice := []rune(s)
	for i := 0; i < len(strSlice); i++ {
		if strSlice[i] < 65 || strSlice[i] > 90 {
			return false
		}
	}
	return true
}
