package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	name := []rune(argument[0])
	for i := 2; i < len(name); i++ {
		z01.PrintRune(rune(name[i]))
	}

	z01.PrintRune('\n')
}
