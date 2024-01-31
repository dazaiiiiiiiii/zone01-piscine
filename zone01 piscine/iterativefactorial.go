package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 31 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	}
	i := nb * (nb - 1) // 4 * 3
	nb--               // nt = 3
	for nb > 1 {
		nb--
		i = i * nb // 12 * 2
	}
	return i
}
