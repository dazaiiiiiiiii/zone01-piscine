package piscine

func ToLower(s string) string {
	str := []rune(s)
	for i := 0; i < len(s); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}
