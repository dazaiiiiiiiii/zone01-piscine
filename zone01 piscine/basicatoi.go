package piscine

func BasicAtoi(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result = result * 10
		result = result + int(s[i]-48)
	}
	return result
}
