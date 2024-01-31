package piscine

func atoi(s string) int {
	result := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if i == 0 && (s[i] == '-') {
			sign = -sign
		} else if s[i] >= '0' && s[i] <= '9' {
			result = result * 10
			result = result + int(s[i]-48)
		}
	}

	return result * sign
}

func TrimAtoi(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			count++
		}
	}
	str := make([]byte, count)
	j := 0
	i := 0
	for ; i < len(s); i++ {
		if (s[i] == '-') || (s[i] >= '0' && s[i] <= '9') {
			str = append([]byte{s[i]}, str...)
			j++
			break
		}
	}
	for k := i + 1; k < len(s); k++ {
		if s[k] >= '0' && s[k] <= '9' {
			str[j] = s[k]
			j++
		}
	}
	strin := string(str)

	return atoi(strin)
}
