package piscine

func JumpOver(str string) string {
	if len(str) < 3 || len(str) < 1 {
		return "\n"
	}
	s := ""
	j := 0
	for i := 2; i < len(str); i += 3 {
		s = s + string(str[i])
		j++
	}
	return s + "\n"
}
