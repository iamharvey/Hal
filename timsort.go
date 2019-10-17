package Hal

/*
TimSort sorts a slice using Tim Sort algorithm. Tim Sort is a sorting algorithm based on Insert Sort and Merge Sort
There are three main steps:
- it first divides a slice into a set of subsequences (known as Run)
- then, it sorts each run using Insert Sort
- at the end, merge those runs using Merge Sort

If the slice size is less than the number of runs, the slice can be sorted just using Insert Sort.
The number of runs varies from 32 to 64, depending on the size of the array. The merge function performs well when
the size of the sequences are powers of 2. For small slice, e.g., less than 100, Insert Sort is recommended over Tim
Sort as the former is much simpler.
 */
func TimSort(a []int, runs int) {
	n := len(a)
	if runs > n {
		// the number of runs should smaller than the slice size, set to default value (32)"
		runs = 32
	} else {
		if n <= 150 {
			runs = n
		} else {
			if runs < 32 || runs > 64 {
				// the number of runs should range from 32 to 64, set to default value (32)"
				runs = 32
			} else {
				if (runs & (runs - 1)) == 0 {
					// the number of runs should be the power of 2, set to default value (32)"
					runs = 32
				}
			}
		}
	}

	// sort each subsequences using Insert Sort
	for start := 0; start < n; start += runs {
		insertSort(a, start, min(start + runs - 1, n - 1))
	}

	// merge runs from the size of n, then 2n, 4n ...

	for size := runs; size < n; size *=2 {
		// for each merge: split the array to two parts and merge
		for start := 0; start < n; start += 2 * size {
			mid := start + size - 1
			end := min(start + 2 * size - 1, n - 1)
			merge(a, start, mid, end)
		}
	}

	return
}

// insertSort sorts a subsequence using Insert Sort
func insertSort(a []int, start, end int) {
	for i := start + 1; i <= end; i ++ {
		curr := a[i]
		j := i - 1
		for j >= start && curr < a[j] {
			a[j + 1] = a[j]
			j --
		}
		a[j +1 ] = curr
	}

	return
}

// merge merges sorted runs
func merge(a []int, start, mid, end int) {

	n1 := mid - start + 1
	n2 := end - mid
	left := make([]int, n1)
	right := make([]int, n2)

	// separate original one into two
	for x := 0; x < n1; x ++ {
		left[x] = a[start + x]
	}

	for x := 0; x < n2; x ++ {
		right[x] = a[mid + 1 + x]
	}

	k := start
	i := 0
	j := 0

	// merge the above two parts in a larger sub slice
	for i < n1 && j < n2  {
		if left[i] <= right[j] {
			a[k] = left[i]
			i ++
		} else {
			a[k] = right[j]
			j ++
		}
		k ++
	}

	// copy the remaining values of the left part
	for i < n1 {
		a[k] = left[i]
		k ++
		i ++
	}

	// copy the remaining values of the right part
	for j < n2 {
		a[k] = right[j]
		k ++
		j ++
	}

	return
}

// min find the smaller out of two given values
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}