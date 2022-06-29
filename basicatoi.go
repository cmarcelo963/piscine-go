package piscine

// 4589
func BasicAtoi(s string) int {
	//			['-', '5', '4', '3', '2', '1']
	aString := []rune(s)
	isNegative := false
	index := 0
	if aString[0] == '-' {
		isNegative = true
		index = 1
	}

	multiplier := 1
	result := 0
	for a := index; a < len(s); a++ {
		multiplier = int(aString[a] - 48)
		for b := a; b < len(s)-1; b++ {
			multiplier *= 10
		}
		result += multiplier
	}
	if isNegative == true {
		result *= -1
	}
	return result
}
