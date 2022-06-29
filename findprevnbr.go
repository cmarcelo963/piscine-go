package piscine

func FindPrevPrime(nb int) int {
	for a := nb - 1; a > 1; a-- {
		if isPrime(a) {
			return a
		}
	}
	return 0
}

func isPrime(nb int) bool {
	prime := true

	for a := 2; a < nb; a++ {
		if (nb % a) == 0 {
			prime = false
		}
	}
	return prime
}
