package piscine

func NRune(s string, n int) rune {
	a := []rune(s)

	if n > 0 && n < len(a) {
		return a[n-1]
	}
	return 0
}
