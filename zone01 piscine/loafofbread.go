package piscine

func LoafOfBread(str string) string {
	if len(str) <= 0 {
		return "\n"
	}
	var result []byte
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	count := 1
	for i := 0; i < len(str); i++ {
		if count != 6 && str[i] != ' ' {
			result = append(result, str[i])
			count++
		}
		if count == 6 && i < len(str)-3 {
			result = append(result, ' ')
			count = 1
			i++
		}
	}
	return string(result) + "\n"
}
