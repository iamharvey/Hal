package temp



// https://en.wikipedia.org/wiki/Odd%E2%80%93even_sort
func OddEvenSort(a []int) []int {
	n := len(a)
	sorted := false
	for !sorted {
		sorted = true
		for i := 1; i < n - 1; i += 2 {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				sorted = false
			}
		}

		for i := 0; i < n - 1; i += 2 {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				sorted = false
			}
		}
	}

	return a
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




