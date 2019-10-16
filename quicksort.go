package Hal

/*
QuickSort sorts an array using iterative quick sort algorithm.
 */
func QuickSort(a []int, scheme string) {

	lo := 0
	hi := len(a) - 1
	quickSortRecur(a, lo, hi, scheme)

	return
}

// sort sorts the array iteratively
func quickSortRecur(a []int, lo, hi int, scheme string) {
	if lo < hi {
		var pos int
		if scheme == "hoare" {
			pos = hoare(a, lo, hi)
			quickSortRecur(a, lo, pos, scheme)
			quickSortRecur(a, pos + 1, hi, scheme)
		} else {
			pos = lomuto(a, lo, hi)
			quickSortRecur(a, lo, pos - 1, scheme)
			quickSortRecur(a, pos + 1, hi, scheme)
		}
	}
}

/*
lomuto partition does three things:
- takes last element as pivot and places the pivot at its correct place
- make all the numbers to the left of the pivot smaller than it
- make all the numbers to the right of the pivot larger than it
*/
func lomuto(a []int, lo, hi int) int {
	pivot := a[hi]

	i := lo
	for j := lo; j < hi; j ++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i ++
		}
	}

	a[i], a[hi] = a[hi], a[i]
	return i
}

/*
hoare partition is similar to lomuto except that:
- takes medium as the pivot and places the pivot at its correct place
- moving check points from both lo->mid and hi->mid
- theoretically hoare should be more efficient than lomuto as it has less swaps
*/
func hoare(a []int, lo, hi int) int {
	pivot := a[lo + ((hi - lo) / 2)]
	i := lo - 1
	j := hi + 1
	for {
		i = i + 1
		for a[i] < pivot {
			i  = i + 1
		}

		j = j - 1
		for a[j] > pivot {
			j = j - 1
		}

		if i >= j {
			return j
		}

		a[i], a[j] = a[j], a[i]
	}
}
