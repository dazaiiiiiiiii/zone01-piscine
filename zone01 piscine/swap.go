package piscine

func Swap(a *int, b *int) {
	swp := *a
	*a = *b
	*b = swp
}
