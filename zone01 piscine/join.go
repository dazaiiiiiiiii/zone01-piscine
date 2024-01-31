package piscine

func Join(strs []string, sep string) string {
	var str string
	var s string
	n := len(strs) - 1
	for _, char := range strs {
		s = char
		if len(s) > 0 {
			str += char
		}
		if char == strs[n] {
			break
		}
		str += sep
	}
	return str
}
