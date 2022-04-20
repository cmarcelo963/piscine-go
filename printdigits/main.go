package main

import "github.com/01-edu/z01"

func main() {
	var digits string = "0123456789"
	digitsRune := []rune(digits)

	for i := 0; i < 10; i++ {
		z01.PrintRune(digitsRune[i])
	}
	z01.PrintRune('\n')

}
