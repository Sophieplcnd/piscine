package piscine

func Rot14(s string) string {
	rot := []rune(s)
	for i := 0; i < len(rot); i++ {
		if (rot[i] >= 'A' && rot[i] <= 'L') || (rot[i] >= 'a' && rot[i] <= 'l') {
			rot[i] += 14
		} else if (rot[i] >= 'L' && rot[i] <= 'Z') || (rot[i] >= 'l' && rot[i] <= 'z') {
			rot[i] -= 12
		}
	}
	return string(rot)
}
