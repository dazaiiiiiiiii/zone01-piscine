package piscine

func FindNextPrime(nb int) int {
	if nb > 2147483647 {
		return 2147483647
	}
	if nb <= 1 {
		return 2
	}

	for i := 2; i <= nb/2; i++ { // nb = 6 3
		if nb%i == 0 { //  8/2
			return FindNextPrime(nb + 1)
		}
	}

	return nb
}
