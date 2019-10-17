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

var arrN = 100

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

//
//func BenchmarkGnomeSort2(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := generateRandomNumber(arrN)
//		GnomeSort2(A)
//	}
//}
//

func BenchmarkTimSort32(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := RandInt(arrN)
		TimSort(A, 32)
	}
}

func BenchmarkTimSort64(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := RandInt(arrN)
		TimSort(A, 64)
	}
}

func BenchmarkTimSort128(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := RandInt(arrN)
		TimSort(A, 128)
	}
}


func BenchmarkQuickSort_Hoare(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := RandInt(arrN)
		QuickSort(A, "hoare")
	}
}

func BenchmarkQuickSort_Lomuto(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := RandInt(arrN)
		QuickSort(A, "lomuto")
	}
}

func BenchmarkQuickSortGolang(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := RandInt(arrN)
		sort.Ints(A)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := RandInt(arrN)
		InsertSort(A)
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
