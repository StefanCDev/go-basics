package piscine

func SplitWhiteSpaces(s string) []string {
	heldpos := 0
	runesl := []rune(s)
	var sep []string
	for i := 1; i < len(runesl); i++ {
		if runesl[i] == ' ' || runesl[i] == '\n' || runesl[i] == '\t' {
			if runesl[i-1] != ' ' && runesl[i-1] != '\n' && runesl[i-1] != '\t' {
				sep = append(sep, string(runesl[heldpos:i]))
			}
			heldpos = i + 1
		} else if i == len(runesl)-1 {
			sep = append(sep, string(runesl[heldpos:i+1]))
		}
	}
	return sep
}
