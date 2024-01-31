package piscine

func MakeRange(min, max int) []int {
	var result []int
	if min >= max {
		return result
	}
	llen := max - min
	result1 := make([]int, llen)
	j := 0
	for i := min; i < max; i++ {
		result1[j] = i
		j++
	}
	return result1
}
