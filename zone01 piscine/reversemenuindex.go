package piscine

func ReverseMenuIndex(menu []string) []string {
	result := make([]string, len(menu))
	j := 0
	for i := len(menu) - 1; i >= 0; i-- {
		result[j] = menu[i]
		j++
	}
	return result
}
