package Hal

import (
	"math"
)

var minRun = 32

func TimSort(a []int) []int {
	n := len(a)
	for i := 0; i < n; i += minRun {
		end := int(math.Min(float64(i + minRun - 1), float64(n - 1)))
		a = InsertSort(a, i, end)
	}

	currSize := minRun
	for currSize < n {
		for i := 0; i < n; i += 2 * currSize {
			mid := int(math.Min(float64(n-1), float64(i + currSize - 1)))
			end := int(math.Min(float64(n-1), float64(mid + currSize)))
			a = merge(a, i, mid, end)
		}
		currSize *= 2
	}

	return a
}

func InsertSort(a []int, start, end int) []int {
	for i := start + 1; i <= end; i ++ {
		temp := a[i]
		j := i - 1
		for j >= start && temp < a[j] {
			a[j+1] = a[j]
			j -= 1
		}
		a[j+1] = temp
	}

	return a
}

func merge(a []int, start, mid, end int) []int {
	if mid == end {
		return a
	}

	first := a[start: mid+1]
	last := a[mid+1 : end+1]
	l1 := mid - start + 1
	l2 := end - mid
	ind1 := 0
	ind2 := 0
	ind := start

	for ind1 < l1 && ind2 < l2 {
		if first[ind1] <= last[ind2] {
			a[ind] = first[ind1]
			ind1 += 1
		} else {
			a[ind] = last[ind2]
			ind2 += 1
		}
		ind += 1
	}

	for ind1 < l1 {
		a[ind] = first[ind1]
		ind1 += 1
		ind += 1
	}

	for ind2 < l2 {
		a[ind] = last[ind2]
		ind2 += 1
		ind += 1
	}

	return a
}