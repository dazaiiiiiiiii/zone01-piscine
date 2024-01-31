package piscine

import "github.com/01-edu/z01"

func Printanum(n int) {
	str := "-9223372036854775808"
	for i := 0; i < len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}
}

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		Printanum(n)
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n <= 9 {
		z01.PrintRune(rune(n + '0'))
	} else {
		PrintNbr(n / 10)
		z01.PrintRune(rune(n%10 + '0'))
	}
}
