package Hal

import (
	"sort"
	"testing"
)

/*
<= 3000: quicksort and quicksortmy consider a better one
3000 - 5000: quicksort and quicksortmy are good
5000 - 10000: quicksortmy2 consistently the best
>= 10000 ( - 100000000): quicksortmy2 consistently the best, quicksort google becomes the second
 */

var arrN = 1000

func BenchmarkGnomeSort(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		GnomeSort(A)
	}
}

func BenchmarkGnomeSort2(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		GnomeSort2(A)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		QuickSort(A)
	}
}

func BenchmarkQuickSortMy(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		QuickSortMy(A, 0, len(A) - 1)
	}
}

func BenchmarkQuickSortMy2(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		QuickSortMy2(A, 0, len(A) - 1)
		// fmt.Println(A)
	}
}

func BenchmarkQuickSortGolang(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		sort.Ints(A)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		BubbleSort(A)
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		BubbleSort2(A)
	}
}

func BenchmarkCocktailShakerSort(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		CocktailShakerSort(A)
	}
}
