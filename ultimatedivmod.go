package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a / *b
	x := *a

	*a = c
	*b = x % *b
}
