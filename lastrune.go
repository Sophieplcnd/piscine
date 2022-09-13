package piscine

func LastRune(s string) rune {
	x := []rune(s)
	length := 0
	for i := range s {
		length = i + 1
	}
	return x[length-1]
}
