package piscine

func Sqrt(nb int) int {
	for a := 2; a < nb; a++ {
		for b := a; b <= nb; b *= b {
			if b == nb {
				return a
			}
		}
	}
	return 0
}
