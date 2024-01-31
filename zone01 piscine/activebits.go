package piscine

func convert(n int) int {
	count := 0
	base := "01"
	nb := n / len(base)
	mod := n % len(base)
	if base[mod] == 49 {
		count++
	}
	for i := 0; nb > 0; i++ {
		mod = nb % len(base)
		nb = nb / len(base)
		if base[mod] == 49 {
			count++
		}
	}
	return count
}

func ActiveBits(n int) int {
	return convert(n)
}
