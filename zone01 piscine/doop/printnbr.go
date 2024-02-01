package main

import (
	"os"
)

func PrintNbr(n int) {
	if n < 0 {
		os.Stdout.WriteString("-")
		n = -n
	}
	printDigit(n)
}

func printDigit(n int) {
	if n/10 != 0 {
		printDigit(n / 10)
		os.Stdout.WriteString(string(byte(n%10 + '0')))
	} else {
		os.Stdout.WriteString(string(byte(n + '0')))
	}
}
