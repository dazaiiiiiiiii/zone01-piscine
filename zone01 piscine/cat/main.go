package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		file, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			for _, c := range "ERROR: " + err.Error() {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
		for _, c := range file {
			z01.PrintRune(rune(c))
		}
	} else {
		for i := 1; i < len(args); i++ {
			file, err := os.Open(args[i])
			if err != nil {
				for _, c := range "ERROR: " + err.Error() {
					z01.PrintRune(c)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			}
			defer file.Close()
			b, err := ioutil.ReadAll(file)
			if err != nil {
				for _, c := range "ERROR: " + err.Error() {
					z01.PrintRune(c)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			}
			for _, c := range b {
				z01.PrintRune(rune(c))
			}
		}
	}
}
