package Hal

import "math/rand"

// RandInt generates a slice of random integers
func RandInt(size int) []int {
	rs := make([]int, size, size)
	for i := 0; i < size; i++ {
		rs[i] = rand.Intn(2 * size)
	}

	return rs
}
