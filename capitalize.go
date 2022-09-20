package piscine

func Capitalize(s string) string {
	aString := ""
	letter := []rune(s)
	for i := 0; i < len(letter); i++ {
		a := string(letter[i])
		// if 'i' is not at index 0 and letter before i is an alphabet, change it to lowercase
		if i != 0 && IsAlpha(string(letter[i-1])) {
			aString += ToLower(a)
			// if letter before 'i' not an alphanumerical, change it to uppercase
		} else {
			aString += ToUpper(a)
		}
	}
	return aString
}
