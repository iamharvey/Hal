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

//func BenchmarkOddEvenSort(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := generateRandomNumber(arrN)
//		A = OddEvenSort(A)
//	}
//}
//
//func BenchmarkBTreeSort(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := generateRandomNumber(arrN)
//		A = BTreeSort(A)
//	}
//}
//
//func BenchmarkRadixSort(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := generateRandomNumber(arrN)
//		A = RadixSort(A)
//	}
//}
//
//func BenchmarkHeapSort(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := generateRandomNumber(arrN)
//		A = HeapSort(A)
//	}
//}
//
//func BenchmarkTimSort(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := generateRandomNumber(arrN)
//		A = TimSort(A)
//	}
//}
//
//func BenchmarkGnomeSort2(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := generateRandomNumber(arrN)
//		GnomeSort2(A)
//	}
//}
//
func BenchmarkQuickSort_Hoare(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		QuickSort(A, "hoare")
	}
}

func BenchmarkQuickSort_Lomuto(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		QuickSort(A, "lomuto")
	}
}

func BenchmarkQuickSortGolang(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := generateRandomNumber(arrN)
		sort.Ints(A)
	}
}

//func BenchmarkBubbleSort(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := generateRandomNumber(arrN)
//		BubbleSort(A)
//	}
//}
//
//func BenchmarkCocktailShakerSort(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := generateRandomNumber(arrN)
//		CocktailShakerSort(A)
//	}
//}
