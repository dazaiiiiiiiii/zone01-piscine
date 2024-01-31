package piscine

func Abort(a, b, c, d, e int) int {
	sort := []int{a, b, c, d, e}
	for i := 0; i < len(sort); i++ {
		for j := i + 1; j < len(sort); j++ {
			if sort[i] > sort[j] {
				sort[i], sort[j] = sort[j], sort[i]
			}
		}
	}
	i := len(sort) / 2
	return sort[i]
}
