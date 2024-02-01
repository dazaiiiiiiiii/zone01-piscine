package piscine

// atoibase but need the checkerbase function to check if the base is valid

func AtoiBase(s string, base string) int {
	n := 0
	nb := 1
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(base); j++ {
			if s[i] == base[j] {
				n = n + (j * nb)
				nb = nb * len(base)
			}
		}
	}
	return n
}
