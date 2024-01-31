package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	if len(arg) > 1 {
		for i := len(arg) - 1; i > 0; i-- {
			for j := 0; j < len(arg[i]); j++ {
				z01.PrintRune(rune(arg[i][j]))
			}
			z01.PrintRune('\n')
		}
	}
}
