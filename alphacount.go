package piscine

func AlphaCount(s string) int {
	counter := 0
	letter := []rune(s)

	for i := 0; i < len(letter); i++ {
		if (letter[i] >= 65 && letter[i] <= 90) || (letter[i] >= 97 && letter[i] <= 122) {
			counter++
		}
	}
	return counter
}
