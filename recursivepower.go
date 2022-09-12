package piscine

func RecursivePower(nb int, power int) int {
	result := nb
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	for i := 1; i < power; i++ {
		result = nb * result
	}

	return result
}