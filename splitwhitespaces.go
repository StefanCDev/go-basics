package piscine

func SplitWhiteSpaces(s string) []string {
	runeSlice := []rune(s)
	var answerSlice []string
	var runeSlice2 []rune
	for i := 0; i < len(runeSlice); i++ {
		if runeSlice[i] != ' ' && i == len(runeSlice)-1 {
			runeSlice2 = append(runeSlice2, runeSlice[i])
			answerSlice = append(answerSlice, string(runeSlice2))
			runeSlice2 = nil
		} else if runeSlice[i] != ' ' {
			runeSlice2 = append(runeSlice2, runeSlice[i])
		} else {
			if runeSlice != nil {
				answerSlice = append(answerSlice, string(runeSlice2))
				runeSlice2 = nil
			}
		}
		return answerSlice
	}
}
