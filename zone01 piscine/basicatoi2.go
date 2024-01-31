package piscine

func BasicAtoi2(s string) int {
	result := 0
	for j := 0; j < len(s); j++ {
		if !(s[j] >= '0' && s[j] <= '9') {
			return 0
		}
	}
	i := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		result = result * 10
		result = result + int(s[i]-48)
		i++
	}
	return result
}
