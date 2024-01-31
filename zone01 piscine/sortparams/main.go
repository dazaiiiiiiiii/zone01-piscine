package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	if len(arg) > 1 {
		for i := 1; i < len(arg); i++ {
			for j := i + 1; j < len(arg); j++ {
				if arg[i] > arg[j] {
					arg[i], arg[j] = arg[j], arg[i]
				}
			}
			for k := 0; k < len(arg[i]); k++ {
				z01.PrintRune(rune(arg[i][k]))
			}
			z01.PrintRune('\n')
		}
	}
}
