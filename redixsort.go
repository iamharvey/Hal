package Hal

//https://www.geeksforgeeks.org/radix-sort/
func RadixSort(a []int) []int {
	n := len(a)
	m := getMax(a, n)

	for exp := 1; m / exp > 0; exp *= 10 {
		a = countSort(a, n, exp)
	}
	return a
}

func getMax(a []int, n int) int {
	mx := a[0]
	for i := 0; i < n; i ++ {
		if a[i] > mx {
			mx = a[i]
		}
	}
	return mx
}

func countSort(a []int, n, exp int) []int {
	op := make([]int, n)
	count := make([]int, 10)
	var i int

	for i := 0; i < n; i ++ {
		count[(a[i] / exp) % 10] ++
	}

	for i = 1; i < 10; i ++ {
		count[i] += count[i - 1]
	}

	for i = n-1; i >= 0; i -- {
		op[count[(a[i] / exp)%10] - 1] = a[i]
		count[(a[i]/ exp) %10] --
	}

	return op
}