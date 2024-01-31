package piscine

func AlphaCount(s string) int {
	i := 0
	count := 0
	for ; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			count++
		}
	}
	return count
}
