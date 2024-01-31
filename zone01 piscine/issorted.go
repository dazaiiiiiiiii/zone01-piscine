package piscine

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	count := 0
	count1 := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			count++
		}
		if f(a[i], a[i+1]) > 0 {
			count1++
		} else if f(a[i], a[i+1]) == 0 {
			return true
		}
	}
	if count1 == len(a)-1 || count == len(a)-1 {
		return true
	} else {
		return false
	}
}
