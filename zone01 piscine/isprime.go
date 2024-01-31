package piscine

func IsPrime(nb int) bool {
	if nb < 0 || nb == 1 || nb == 0 {
		return false
	}
	for i := nb - 1; i > 1; i-- { // i = 5
		if nb%i == 0 { // 5 / 4
			return false
		}
	}
	return true
}
