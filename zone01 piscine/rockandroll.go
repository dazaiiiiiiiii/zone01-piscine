package piscine

func RockAndRoll(n int) string {
	s1 := "error: number is negative\n"
	s2 := "error: non divisible\n"
	if n < 0 {
		return s1
	} else if n%3 == 0 && n%2 == 0 {
		return "rock and roll\n"
	} else if n%2 == 0 {
		return "rock\n"
	} else if n%3 == 0 {
		return "roll\n"
	}
	return s2
}
