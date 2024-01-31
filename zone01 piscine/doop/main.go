package main

import (
	"os"
)

func IsNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if !((s[i] >= '0' && s[i] <= '9') || (s[0] == '-' || s[0] == '+')) {
			return false
		}
	}
	return true
}

func Itoa(nb int) {
	if nb < 0 {
		nb = -nb
		digite := 0
		temp := nb
		for temp > 0 {
			temp /= 10
			digite++
		}
		result := make([]byte, digite+1)
		for i := len(result) - 1; i >= 0; i-- {
			result[i] = byte(nb%10) + '0'
			nb /= 10
		}
		result[0] = '-'
		printt(string(result))
		return
	}
	if nb == 0 {
		result := make([]byte, nb)
		result = append(result, '0')
		printt(string(result))
		return
	}
	digite := 0
	temp := nb
	for temp > 0 {
		temp /= 10
		digite++
	}
	result := make([]byte, digite)
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = byte(nb%10) + '0'
		nb /= 10
	}

	printt(string(result))
}

func printt(s string) {
	str := []byte(s)

	os.Stdout.Write(str)
	os.Stdout.WriteString("\n")
}

func calcul(nbs []int, op rune) {
	result := 0
	if op == '+' {
		result = nbs[0] + nbs[2]
	} else if op == '-' {
		result = nbs[0] - nbs[2]
	} else if op == '*' {
		result = nbs[0] * nbs[2]
	} else if op == '/' {
		result = nbs[0] / nbs[2]
	} else if op == '%' {
		result = nbs[0] % nbs[2]
	} else {
		return
	}
	Itoa(result)
}

func atoi(s string) int {
	result := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] == '-' || s[i] == '+' {
			if s[i] == '-' {
				sign = -sign
			}
		} else if s[i] >= '0' && s[i] <= '9' {
			result *= 10
			result += int(s[i] - '0')
		} else {
			return int(s[i])
		}
	}
	if result*sign > 9223372036854775807 {
		return 0
	} else if result*sign < -9223372036854775808 {
		return 0
	}

	return result * sign
}

func main() {
	arg := os.Args
	var res []int
	if len(arg) > 3 {
		if arg[1] == "9223372036854775807" || arg[3] == "9223372036854775807" {
			if arg[2] == "+" || arg[2] == "*" {
				return
			}
		} else if arg[1] == "-9223372036854775808" || arg[3] == "-9223372036854775808" {
			if arg[2] == "-" || arg[2] == "*" {
				return
			}
		} else if arg[2] == "/" && arg[3] == "0" {
			os.Stdout.WriteString("No division by 0\n")
			return
		} else if arg[2] == "%" && arg[3] == "0" {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}

		if IsNumeric(arg[1]) && IsNumeric(arg[3]) {

			for i := 1; i < len(arg); i++ {
				res = append(res, atoi(arg[i]))
			}

			calcul(res, rune(arg[2][0]))
		}

	}
}
