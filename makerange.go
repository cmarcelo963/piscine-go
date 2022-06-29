package piscine

func MakeRange(min, max int) []int {
	var result []int
	if min > max {
		result = nil
		return result
	}

	result = make([]int, min, max)
	for a := min; a < max; a++ {
		counter := 0
		result[counter] = a
		counter++
	}
	return result
}
