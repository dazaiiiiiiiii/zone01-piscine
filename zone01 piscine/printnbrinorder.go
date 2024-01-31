package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	digite := 0
	temp := n
	for temp > 0 {
		temp /= 10
		digite++
	}
	result := make([]rune, digite)
	for i := 0; i < len(result); i++ {
		result[i] = rune(n%10 + '0')
		n /= 10
	}
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	for i := 0; i < len(result); i++ {
		z01.PrintRune(rune(result[i]))
	}
}
