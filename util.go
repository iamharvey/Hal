package Hal

import (
	"math/rand"
)

// RandInt generates a slice of random integers
func RandInt(size int) []int {
	rs := make([]int, size, size)
	for i := 0; i < size; i++ {
		rs[i] = rand.Intn(2 * size)
	}

	return rs
}

// RandIntUnique generate a slice of random integers with no duplication
func RandIntUnique(size int) []int {
	rHolder := make(map[int]bool, size)
	rs := make([]int, size, size)
	for i := 0; i < size; i++ {
		var r int
		for true {
			r = rand.Intn(size)
			if ! rHolder[r] {
				rHolder[r] = true
				break
			}
		}
		rs[i] = r
	}

	return rs
}