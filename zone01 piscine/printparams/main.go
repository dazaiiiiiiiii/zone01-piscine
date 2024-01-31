package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	if len(arg) > 1 {
		for i := 1; i < len(arg); i++ {
			for j := 0; j < len(arg[i]); j++ {
				z01.PrintRune(rune(arg[i][j]))
			}
			z01.PrintRune('\n')
		}
	}
}
