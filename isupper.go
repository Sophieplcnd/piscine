package piscine

func IsUpper(s string) bool {
	letter := []rune(s)
	for i := 0; i < len(letter); i++ {
		if s[i] < 65 || s[i] > 90 {
			return false
		}
	}
	return false
}
