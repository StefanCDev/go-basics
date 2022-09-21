package main

import (
	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func StripDigits(n int) []int {
	if n < 10 {
		return []int{n}
	} else {
		return append(StripDigits(n/10), []int{n % 10}...)
	}
}

func setPoint(ptr *[2]int) {
	ptrArr := *ptr
	ptrArr[0] = 42
	ptrArr[1] = 21
	*ptr = ptrArr
}

func main() {
	points := [2]int{}
	numArr := make([]int, 0)

	setPoint(&points)

	numArr = StripDigits(points[0])

	printStr("x = ")
	for _, r := range numArr {
		z01.PrintRune(rune(r + 48))
	}
	printStr(", y = ")
	numArr = StripDigits(points[1])
	for _, r := range numArr {
		z01.PrintRune(rune(r + 48))
	}
	z01.PrintRune('\n')
}
