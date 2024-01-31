package piscine

func findsqrt(n int) int {
	count := 0
	for i := 1; n >= 0; i += 2 {
		if n == 0 {
			break
		}
		n = n - i
		count++
	}
	if n < 0 {
		return 0
	}
	return count
}

func Sqrt(nb int) int {
	return findsqrt(nb)
}
