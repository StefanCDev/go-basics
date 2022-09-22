package main

import (
	"os"
	"strconv"
)

func main() {
	for len(os.Args) == 2 {
		numbers, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		newnum := numbers
		counts := 0
		for numbers != 1 {
			if newnum%2 != 0 {
			} else {
				newnum = newnum / 2
			}
			counts++
		}
		var x int = 2 ^ counts
		if x == numbers {
			return
		}
	}
}
