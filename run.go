package Hal

import "math/rand"

func generateRandomNumber(size int) []int {
	rand_number := make([]int, size, size)
	for i := 0; i < size; i++ {
		rand_number[i] = rand.Intn(100)
	}

	return rand_number
}

// https://en.wikipedia.org/wiki/Gnome_sort
func GnomeSort(arr []int) []int {
	n := len(arr)
	p := 0
	for p < n {
		if p == 0 || (arr[p] >= arr[p-1]) {
			p = p + 1
		} else {
			arr[p], arr[p - 1] = arr[p-1], arr[p]
			p = p - 1
		}
	}

	return arr
}

// https://en.wikipedia.org/wiki/Gnome_sort
func GnomeSort2(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i ++ {
		j := i
		for j > 0 && arr[j - 1] > arr[j] {
			arr[j - 1], arr[j] = arr[j], arr[j - 1]
			j = j - 1
		}
	}
	return arr
}

func BubbleSort(arr []int) []int {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n-1; i ++ {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1]  = arr[i-1], arr[i]
				swapped = true
			}
		}

		n = n - 1
	}

	return arr
}


func BubbleSort2(arr []int) []int {
	n := len(arr)
	x := n
	for x > 1 {
		x = 0
		for i := 1; i < n; i ++ {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1]  = arr[i-1], arr[i]
				x = i
			}
		}
	}

	return arr
}

func CocktailShakerSort(arr []int) []int {
	n := len(arr)
	start := 1
	end := n - 1
	for start <= end {
		newStart := end
		newEnd := start
		for i := start; i < end; i ++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1]  = arr[i+1], arr[i]
				newEnd = i
			}
		}

		end = newEnd - 1

		for i := end; i > start; i -- {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1]  = arr[i+1], arr[i]
				newStart = i
			}
		}

		start = newStart + 1
	}

	return arr
}

func QuickSortMy(arr []int, lo, hi int) []int {

	if lo < hi {
		p := Partition(arr, lo, hi)
		arr = QuickSortMy(arr, lo, p-1)
		arr = QuickSortMy(arr, p+1, hi)
	}

	return arr
}

func Partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	i := lo
	for j := lo; j < hi; j ++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			 i = i + 1
		}
	}

	arr[i], arr[hi] = arr[hi], arr[i]
	return i
}

func QuickSortMy2(arr []int, lo, hi int) []int {

	if lo < hi {
		p := Partition2(arr, lo, hi)
		arr = QuickSortMy2(arr, lo, p)
		arr = QuickSortMy2(arr, p+1, hi)
	}

	return arr
}

func Partition2(arr []int, lo, hi int) int {
	pivot := arr[lo + ((hi - lo) / 2)]
	i := lo - 1
	j := hi + 1
	for {
		i = i + 1
		for arr[i] < pivot {
			i  = i + 1
		}

		j = j - 1
		for arr[j] > pivot {
			j = j - 1
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}

func QuickSort(arr []int) []int {
	// clone arr to keep immutability
	newArr := make([]int, len(arr))

	for i, v := range arr {
		newArr[i] = v
	}

	// call recursive funciton with initial values
	recursiveSort(newArr, 0, len(newArr)-1)

	// at this point newArr is sorted
	return newArr
}

func recursiveSort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	// Iterate sub array to find values less than pivot
	//   and move them to the beginning of the array
	//   keeping splitIndex denoting less-value array size
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			if splitIndex != i {
				temp := arr[splitIndex]

				arr[splitIndex] = arr[i]
				arr[i] = temp
			}

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	recursiveSort(arr, start, splitIndex-1)
	recursiveSort(arr, splitIndex+1, end)
}


