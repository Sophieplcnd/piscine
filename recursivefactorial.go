package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 || nb == 0 { // base case
		return 1
	}

	return nb * RecursiveFactorial(nb-1) // recursive case
}
