package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	sign := 1

	if nbr < 0 {
		sign = -sign
	}
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if (base[j] == base[i]) || (base[i] == '-' || base[i] == '+') {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	var result1 []byte
	lbase := len(base)
	n := nbr / lbase
	mod := nbr % lbase
	if nbr < 0 {
		mod = -mod
		n = -n
	}
	result1 = append(result1, base[mod])
	for i := 0; n > 0; i++ {
		mod = n % lbase
		n = n / lbase
		if n < 0 {
			mod = -mod
			n = -n
		}
		result1 = append(result1, base[mod])
	}
	if sign < 0 {
		z01.PrintRune('-')
	}
	result2 := string(result1)
	for i := len(result2) - 1; i >= 0; i-- {
		z01.PrintRune(rune(result2[i]))
	}
}
