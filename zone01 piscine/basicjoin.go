package piscine

func BasicJoin(elems []string) string {
	var str string

	for _, char := range elems {
		str += char
	}
	return str
}
