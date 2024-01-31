package piscine

func NRune(s string, n int) rune {
	if n > len(s) || n < 0 {
		return rune(0)
	}
	i := 1
	for _, char := range s {
		if i == n {
			return char
		}
		i++
	}
	return rune(0)
}
