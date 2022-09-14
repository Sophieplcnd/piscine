package piscine

func IsAplha(s string) bool {
	letter := []rune(s)
	for i := 0; i < len(letter); i++ {
		if (s[i] < 65 || s[i] > 90) && (s[i] < 97 || s[i] > 122) && (s[i] < 48 || s[i] > 57) {
			return false
		}
	}
	return true
}
fwd