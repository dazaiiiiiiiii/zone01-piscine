package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if count != i {
				result = append(result, s[count:i])
			}
			count = i + 1
		}
	}
	if count < len(s) {
		result = append(result, s[count:])
	}
	return result
}
