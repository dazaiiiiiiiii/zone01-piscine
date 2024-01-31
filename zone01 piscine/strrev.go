package piscine

func StrRev(s string) string {
	str := []byte(s)

	i := 0
	j := len(str) - 1
	for i < j {
		swp := str[i]
		str[i] = str[j]
		str[j] = swp
		i++
		j--
	}
	return string(str)
}
