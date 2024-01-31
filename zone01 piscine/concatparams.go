package piscine

func ConcatParams(args []string) string {
	var result []byte
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			result = append(result, args[i][j])
		}
		if i < len(args)-1 {
			result = append(result, '\n')
		}
	}
	return (string(result))
}
