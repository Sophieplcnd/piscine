package piscine

func TrimAtoi(s string) int {
	runes := []rune(s)
	aString := ""
	check := false
	for i := 0; i < len(runes); i++ {
		if runes[i] == 45 && !check {
			aString = "-"
		} else if runes[i] >= 48 && runes[i] <= 57 {
			check = true
			aString += string(runes[i])
		}
	}
	return "Atoi"(aString)
}
