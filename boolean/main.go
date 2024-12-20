package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)
	for _, s := range arrayStr {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}

func even(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false
}

func isEven(nbr int) bool {
	if even(nbr) == true {
		return true
	}
	return false
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	lenOfArg := len(os.Args[1:])
	if isEven(lenOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
