package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args // insperation flex
	if len(arg) > 1 {
		for i := 1; i < len(arg); i++ {
			n := 0
			for j := 0; j < len(arg[i]); j++ {
				n *= 10
				n += int(arg[i][j] - '0')
			}
			if arg[1] == "--upper" {
				z01.PrintRune(rune(n + 64))
			} else if n >= 1 && n <= 26 {
				z01.PrintRune(rune(n + 96))
			} else {
				z01.PrintRune(' ')
			}

		}
		z01.PrintRune('\n')
	}
}
