package Hal

/*
HeapSort sorts a slice using Heap Sort algorithm. Heap Sort is based on Binary Heap data structure.
A binary heap is a complete binary tree where value in a parent node is larger (or smaller) than those in its child
nodes. If 'larger' approach is applied, the heap is called max heap, otherwise it is called min heap. For a full binary
tree, if a parent node index is i, then its child node indices are 2i +1 and 2i + 2. There are three steps:
- build a max (or min) heap over the target slice
- replace the root (which is the largest in max heap) with the last item of the heap followed by reducing the size of
heap by 1
- heapify the root
- repeat the above steps while the heap size > 1

asc = true, ascending order
asc = false, descending order
 */
func HeapSort(a []int, asc bool) {
	n := len(a)

	// build max/min heap
	for i := n/2 - 1; i >= 0; i-- {
		if asc {
			heapifyMax(a, n, i)
		} else {
			heapifyMin(a, n, i)
		}
	}

	for i := n - 1; i >= 0; i-- {
		// replace root with current one
		a[i], a[0] = a[0], a[i]

		// reduce the heap size and build max/min heap upon the new root
		if asc {
			heapifyMax(a, i, 0)
		} else {
			heapifyMin(a, i, 0)
		}
	}

	return
}

/*
heapify here is the procedure of building max heap on node. It is applicable only if the children of the node are
heapified. Thus, it is bottom-up approach.
 */
func heapifyMax(a []int, n, i int) {

	largest := i
	left := 2 * i + 1
	right := 2 * i + 2

	// if the left child is large than root, set largest index to left child index
	if left < n && a[left] >  a[largest] {
		largest = left
	}

	// if the right child is large than root, set largest index to right child index
	if right < n && a[right] > a[largest] {
		largest = right
	}

	// if the largest is not root, keep on the heap reduction and heapify procedure
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		heapifyMax(a, n, largest)
	}

	return
}

/*
heapify here is the procedure of building max heap on node. It is applicable only if the children of the node are
heapified. Thus, it is bottom-up approach.
*/
func heapifyMin(a []int, n, i int) {

	smallest := i
	left := 2 * i + 1
	right := 2 * i + 2

	// if the left child is large than root, set largest index to left child index
	if left < n && a[left] <  a[smallest] {
		smallest = left
	}

	// if the right child is large than root, set largest index to right child index
	if right < n && a[right] < a[smallest] {
		smallest = right
	}

	// if the largest is not root, keep on the heap reduction and heapify procedure
	if smallest != i {
		a[i], a[smallest] = a[smallest], a[i]
		heapifyMin(a, n, smallest)
	}

	return
}