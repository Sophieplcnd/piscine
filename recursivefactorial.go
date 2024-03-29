package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 9999 { // base case
		return 0
	} else if nb == 0 {
		return 1
	}

	if nb == 1 {
		return 1
	}

	return nb * RecursiveFactorial(nb-1) // recursive case
}
