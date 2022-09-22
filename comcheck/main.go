package main

import (
	"fmt"
	"os"
)

func main() {
	safeWords := []string{"01", "galaxy", "galaxy 01"}
	alert := false

	for i := 1; i < len(os.Args); i++ {
		for j := 0; j < 3; j++ {
			if os.Args[i] == safeWords[j] {
				alert = true
			}
		}
	}
	if alert {
		fmt.Println("Alert!!!")
	}
}
