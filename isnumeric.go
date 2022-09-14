package piscine

func IsNumeric(s string) bool {
	letter := []rune(s)
	for i := 0; i < len(letter); i++ {
		if s[i] < 48 || s[i] > 57 {
			return false
		}
	}
	return true
}
