package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	i := 0
	for ; i < len(a); i++ {
		if i < len(a)-1 && a[i] == a[i+1] {
			i++
		} else if i <= len(a) {
			return a[i]
		}
	}

	return -1
}
