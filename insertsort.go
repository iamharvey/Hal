package Hal

/*
InsertSort sorts a slice using Insert Sort algorithm. There key operation is to place
a next element into previously sorted sequence.
Insert sort works pretty fast for small slice, e.g., 100 or less elements.
 */
func InsertSort(a []int) {
	n := len(a)
	var k, j int
	for i := 1; i < n; i ++ {
		k = a[i]

		for j = i - 1; j >= 0 && a[j] > k; j -- {
			// move one position ahead
			a[j + 1] = a[j]
		}

		a[j + 1] = k
	}
}
