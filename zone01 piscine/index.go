package piscine

func Index(s string, toFind string) int {
	if len(toFind) < 0 || len(s) < 0 {
		return 0
	}
	index := 0
	temp := 0
	k := 1
	for i := 0; i < len(toFind); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] == toFind[i] {
				if temp == 0 {
					temp++
					index = j
				}
				k++
				if len(toFind) < k {
					return index
				}
			}
		}
	}
	return -1
}
