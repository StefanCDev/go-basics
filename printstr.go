package piscine

import (
	"github.com/-1-edu/z01"
)

func PrintStr(s string) {
	l := []rune(s)
	for _, char := range l {
		z01.PrintRune(char)
	}
}
