package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nb int) bool {
	if nb%2 == 0 {
		return true
	}
	return false
}

func isEven(nbr int) bool {
	yes := true
	no := false
	if even(nbr) {
		return yes
	} else {
		return no
	}
}

func main() {
	lengthOfArg := len(os.Args) - 1
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
