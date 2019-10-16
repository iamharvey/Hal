package Hal

// https://www.programiz.com/dsa/heap-sort
func HeapSort(a []int) []int {
	n := len(a)
	for i := n / 2 - 1; i >= 0; i -- {
		a = heapify(a, n, i)
	}

	for i := n-1; i >= 0; i -- {
		a[i], a[0] = a[0], a[i]
		a = heapify(a, i, 0)
	}

	return a
}

func heapify(a []int, n, i int) []int {
	largest := i
	start := 2 * i + 1
	end := 2 * i + 2

	if start < n && a[start] >  a[largest] {
		largest = start
	}

	if end < n && a[end] > a[largest] {
		largest = end
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		a = heapify(a, n, largest)
	}

	return a

}