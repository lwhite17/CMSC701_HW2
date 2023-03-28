package run_test

import (
	"testing"

	"github.com/lwhite17/CMSC701_HW2/run"
)

func TestRun1(t *testing.T) {
	run.RunTask1()
}

func TestRun2(t *testing.T) {
	run.RunTask2()
}

func TestRun3(t *testing.T) {
	run.RunTask3()
}

//
// How much space are we saving for sparse arrays?

func BenchmarkSlice1000_1(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 1,000 (1 - .01) empty elements
	nel := uint64(1000 - 1000*0.01)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}

func BenchmarkSlice1000_5(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 1,000 (1 - .05) empty elements
	nel := uint64(1000 - 1000*0.05)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}

func BenchmarkSlice1000_10(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 1,000 (1 - .1) empty elements
	nel := uint64(1000 - 1000*0.1)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}

func BenchmarkSlice10000_1(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 10,000 (1 - .01) empty elements
	nel := uint64(10000 - 10000*0.01)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}

func BenchmarkSlice10000_5(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 10,000 (1 - .05) empty elements
	nel := uint64(10000 - 10000*0.05)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}

func BenchmarkSlice10000_10(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 10,000 (1 - .1) empty elements
	nel := uint64(10000 - 10000*0.1)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}

func BenchmarkSlice100000_1(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 100,000 (1 - .01) empty elements
	nel := uint64(100000 - 100000*0.01)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}

func BenchmarkSlice100000_5(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 100,000 (1 - .05) empty elements
	nel := uint64(100000 - 100000*0.05)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}

func BenchmarkSlice100000_10(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 100,000 (1 - .1) empty elements
	nel := uint64(100000 - 100000*0.1)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}

func BenchmarkSlice1000000_1(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 1,000,000 (1 - .01) empty elements
	nel := uint64(1000000 - 1000000*0.01)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}

func BenchmarkSlice1000000_5(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 1,000,000 (1 - .05) empty elements
	nel := uint64(1000000 - 1000000*0.05)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}

func BenchmarkSlice1000000_10(b *testing.B) {
	// Benchmark generating a slice of strings
	// with 1,000,000 (1 - .1) empty elements
	nel := uint64(1000000 - 1000000*0.1)
	for i := 0; i < b.N; i++ {
		_ = make([]string, nel)
	}
}
