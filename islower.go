package piscine

func IsLower(s string) bool {
	strSlice := []rune(s)
	for i := 0; i < len(strSlice); i++ {
		if strSlice[i] < 97 || strSlice[i] > 122 {
			return false
		}
	}
	return true
}
