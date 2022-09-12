package piscine

func IterativeFactorial(nb int) int {
	f := 1 // factorial variable

	if nb < 0 || nb > 20 { // conditional statement testing for errors
		f = 0 // have to return factorial as 0
	} else { // factorial
		for i := 1; i <= nb; i++ {
			f = f * i
		}
	}

	return f
}
