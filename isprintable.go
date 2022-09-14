package piscine

func IsPrintable(s string) bool {
	letter := []rune(s)
	for i := 0; i < len(letter); i++ {
		if s[i] < 32 {
			return false
		}
	}
	return true
}
