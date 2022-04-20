package main

import "github.com/01-edu/z01"

func main() {
	var alpha string = "zyxwvutsrqponmlkjihgfedcba"
	alphaRune := []rune(alpha)
//	alphaRune = ['z', 'y', 'x', 'w', 'v']
	for i := 0; i < 26; i++ {
		z01.PrintRune(alphaRune[i])
	}
	z01.PrintRune('\n')
}