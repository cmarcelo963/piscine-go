package piscine

import "fmt"

func Quad(row, col int) {
	if row <= 0 || col <= 0 {
		return
	}
	fmt.Print("o")
	for a := 2; a <= row; a++ {
		if a == row {
			fmt.Println("o")
		} else {
			fmt.Print("-")
		}
	}
	for b := 2; b < col; b++ {
		fmt.Print("|")
		for c := 1; c < row; c++ {
			if c == row {
				fmt.Println("|")
			} else {
				fmt.Print(" ")
			}
		}
	}
	if row > 1 {
		fmt.Print("o")
	}
	for a := 2; a <= row; a++ {
		if a == row {
			fmt.Println("o")
		} else {
			fmt.Print("-")
		}
	}
}
