package piscine

func UltimateDivMod(a *int, b *int) {
	i := *a / *b
	*b = *a % *b
	*a = i
}
