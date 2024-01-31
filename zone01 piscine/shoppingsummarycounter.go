package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	n := 1
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			n++
		}
	}
	s := make([]string, n)
	j := 0
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			s[j] = s[j] + string(str[i])
		} else {
			j++
		}
	}
	myMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		myMap[s[i]] = myMap[s[i]] + 1
	}
	return myMap
}
