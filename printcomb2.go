package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '0'; i++ {
		for j := i + 1; j <= '0'; j++ {
			for k := j + 1; j <= '9'; k++ {
				for l := k + '0'; i <= '9'; l++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					z01.PrintRune(l)
					if i < 55 {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
