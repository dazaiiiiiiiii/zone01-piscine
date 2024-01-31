package piscine

import "github.com/01-edu/z01"

func printthecomb(arr []rune, begin, n, remain int) {
	if remain == 0 {
		for i := 0; i < n; i++ {
			z01.PrintRune(arr[i])
		}
		if 10-int(arr[0]-'0') != n {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}
	for i := begin; i <= 9; i++ {
		newarr := append(arr, rune(i+'0'))
		printthecomb(newarr, i+1, n, remain-1)
	}
}

func PrintCombN(n int) {
	if n < 0 || n > 10 {
		return
	}
	printthecomb([]rune{}, 0, n, n)
	z01.PrintRune('\n')
}
