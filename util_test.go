package Hal

import (
	"testing"
)

func BenchmarkRandIntUnique(b *testing.B) {
	for n := 0; n < b.N; n ++ {
		RandIntUnique(1000)
	}
}
