package Hal

/*
Sort is a simple wrap-up over fast sorting algorithms.
If the slice size is leq 100, Insert Sort is used;
otherwise quick sort is used. This setup is based on my benchmark results.
 */
func Sort(a []int) {
	n := len(a)
	if n <= 100 {
		InsertSort(a)
	} else {
		QuickSort(a, "lomuto")
	}
}
