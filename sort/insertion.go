package goalgorithms

// InsertionSortSwap sorts an int slice in ascending order by swapping values.
// Worst case time compexity: O(n^2)
// Worst case space compexity: O(n)
func InsertionSortSwap(a []int) {
	for i := 1; i < len(a); i++ {
		k := i
		for k > 0 && a[k] < a[k-1] {
			a[k], a[k-1] = a[k-1], a[k]
			k--
		}
	}
}

// InsertionSortSwapOnce sorts an int slice in ascending order by swapping once in the inner loop.
// Worst case time compexity: O(n^2)
// Worst case space compexity: O(n)
func InsertionSortSwapOnce(a []int) {
	for i := 1; i < len(a); i++ {
		v := a[i]
		k := i
		for k > 0 && v < a[k-1] {
			a[k] = a[k-1]
			k--
		}
		a[k] = v
	}
}

// InsertionSortShift sorts an int slice in ascending order by shifting values.
func InsertionSortShift(a []int) {
	for i := 1; i < len(a); i++ {
		k := i
		temp := a[i]
		for k > 0 && a[i] < a[k-1] {
			k--
		}
		for m := i; m > k; m-- {
			a[m] = a[m-1]
		}
		a[k] = temp
	}
}
