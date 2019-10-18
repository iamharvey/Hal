package Hal

import (
	"testing"
)


var arrN = 100000

//func BenchmarkBTreeSort(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := RandInt(arrN)
//		temp.BTreeSort(A)
//	}
//}
//
//func BenchmarkRadixSort(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := RandInt(arrN)
//		temp.RadixSort(A)
//	}
//}

func BenchmarkHeapSortASC(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := RandInt(arrN)
		HeapSort(A, true)
	}
}

//func BenchmarkHeapSortDES(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := RandInt(arrN)
//		HeapSort(A, false)
//	}
//}
//
//func BenchmarkTimSort32(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := RandInt(arrN)
//		TimSort(A, 32)
//	}
//}
//
//func BenchmarkTimSort64(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := RandInt(arrN)
//		TimSort(A, 64)
//	}
//}
//
//func BenchmarkTimSort128(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := RandInt(arrN)
//		TimSort(A, 128)
//	}
//}
//
//
func BenchmarkQuickSort_Hoare(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		A := RandIntUnique(arrN)
		QuickSort(A, "hoare")
	}
}
//
//func BenchmarkQuickSort_Lomuto(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := RandInt(arrN)
//		QuickSort(A, "lomuto")
//	}
//}
//
//func BenchmarkQuickSortGolang(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := RandInt(arrN)
//		sort.Ints(A)
//	}
//}
//
//func BenchmarkInsertSort(b *testing.B) {
//	for n := 0; n < b.N; n ++ {
//		A := RandInt(arrN)
//		InsertSort(A)
//	}
//}
