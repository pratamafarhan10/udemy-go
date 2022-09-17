package main

import "testing"

func BenchmarkRenameWithoutConcurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		renameWithoutConcurrency([]string{"andre", "taulany", "adul", "future", "drake", "sule", "komeng"})
	}
}

func BenchmarkRenameWithConcurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		renameWithConcurrency([]string{"andre", "taulany", "adul", "future", "drake", "sule", "komeng"})
	}
}
