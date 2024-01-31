package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	j := nb * nb
	for i := 0; i < power-2; i++ {
		j *= nb
	}
	return j
}
