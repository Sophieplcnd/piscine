package piscine

func Abort(a, b, c, d, e int) int {
	medianrange := []int{a, b, c, d, e}       // create the slice medianange that takes 5 integers
	for i := 0; i < len(medianrange)-1; i++ { // first loop
		for j := i + 1; j < len(medianrange); j++ { // start second loop at first loop +1 so it doesnt read the same number twice but instead comapares the first index against the second etc
			if medianrange[i] > medianrange[j] { // if the i index is more than the index of j (this is how you sort)
				medianrange[i], medianrange[j] = medianrange[j], medianrange[i] // then swap index i and j (so the smaller index goes before the larger in the array)
			}
		}
	}
	return medianrange[2] // return the value at the second index as the secod index is where the middle number would be in an array of 5 integers
}
