package piscine

// package main
/*     var FirstRune rune
    for _, c := range s {
        FirstRune = c
        break
} */
func FirstRune(s string) rune {
	arrayStr := []rune(s)
	return arrayStr[0]
}

/* func main() {
    z01.PrintRune(piscine.FirstRune("Hello!"))
    z01.PrintRune(piscine.FirstRune("Salut!"))
    z01.PrintRune(piscine.FirstRune("Ola!"))
    z01.PrintRune('\n')
}
*/
