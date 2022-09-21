package piscine

func Map(f func(int) bool, a []int) []bool {
	boolslice := make([]bool, len(a))
	for i, v := range a {
		boolslice[i] = f(v)
	}
	return boolslice
}
