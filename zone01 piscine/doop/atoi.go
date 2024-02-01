package main

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	var isNegative bool
	arr := []rune(s)
	i := 0
	nb := 0

	if arr[0] == '-' {
		isNegative = true
		i++
	} else if arr[0] == '+' {
		i++
	}

	for ; i < len(s); i++ {
		if arr[i] >= '0' && arr[i] <= '9' {
			nb *= 10
			nb += int(arr[i] - '0')
		} else {
			return 0
		}
	}

	if isNegative {
		return -nb
	}
	return nb
}
