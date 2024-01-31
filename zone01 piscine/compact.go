package piscine

func Compact(ptr *[]string) int {
	slice := make([]string, len(*ptr))
	j := 0
	count := 0
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			slice[j] = (*ptr)[i]
			j++
			count++
		}
	}
	*ptr = slice[:j]
	return count
}
