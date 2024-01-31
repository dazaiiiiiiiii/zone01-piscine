package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var result []byte
	var fivechars []byte
	count := 0
	for i := 0; i < len(str); i++ {
		for i < len(str) && count < 5 {
			for str[i] == ' ' && i < len(str)-1 {
				i++
			}
			if str[i] != ' ' {
				fivechars = append(fivechars, str[i])
			}
			count++
			i++
		}
		count = 0
		for j := 0; j < len(fivechars); j++ {
			result = append(result, fivechars[j])
		}
		if i < len(str)-1 && check(str, i) {
			result = append(result, ' ')
		}

		fivechars = []byte{}
	}

	return string(result) + "\n"
}

func check(s string, i int) bool {
	count := 0
	j := 0
	for ; i < len(s); j++ {
		if s[i] == ' ' {
			count++
		}
		i++
	}
	if count == j {
		return false
	}
	return true
}

