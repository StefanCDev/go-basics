package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for b := 0; b < len(s); b++ {
		z01.PrintRune(rune(s[b]))
	}
}
