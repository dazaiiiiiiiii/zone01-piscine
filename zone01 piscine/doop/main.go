package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	// Error Cases
	if !IsNumeric(args[0]) || !IsNumeric(args[2]) {
		return
	}

	// Limits cases
	if len(args[0]) >= 19 || len(args[2]) >= 19 {
		return
	}

	rst := 0
	a := Atoi(args[0])
	op := args[1]
	b := Atoi(args[2])

	// Divition by 0
	if op == "/" && b == 0 {
		os.Stdout.WriteString("No division by 0\n")
		return
	}
	// Modulo by 0
	if op == "%" && b == 0 {
		os.Stdout.WriteString("No modulo by 0\n")
		return
	}

	// Operations
	switch op {
	case "+":
		rst = add(a, b)
	case "-":
		rst = sub(a, b)
	case "*":
		rst = mul(a, b)
	case "/":
		rst = div(a, b)
	case "%":
		rst = mod(a, b)
	default:
		return
	}

	PrintNbr(rst)
	os.Stdout.WriteString("\n")
}
