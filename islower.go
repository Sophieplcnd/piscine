package piscine

func IsLower(s string) bool {
	letter := []rune(s)
	for i := 0; i < len(letter); i++ {
		if s[i] < 97 || s[i] > 122 {
			return false
		}
	}
	return true
}
