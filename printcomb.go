package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '0'; a <= '7'; a++ {
		for b := 1 + a; b <= '8'; b++ {
			for c := 1 + b; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				if a <= '6' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('\n')
				}
			}
		}
	}
}
