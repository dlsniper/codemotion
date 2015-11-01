package bench1_test

import (
	"testing"

	"github.com/dlsniper/codemotion/code/bench/1"
)

var sorted []int

func TestInsertionSort(t *testing.T) {
	toSort := []int{31, 41, 59, 26, 53, 58, 97, 93, 23, 84}
	sorted := []int{23, 26, 31, 41, 53, 58, 59, 84, 93, 97}
	bench1.InsertionSort(toSort)

	if len(toSort) != len(sorted) {
		t.Logf("len mismatch. expected: %d got: %d", len(sorted), len(toSort))
		t.Fail()
	}

	for idx := range sorted {
		if toSort[idx] != sorted[idx] {
			t.Logf("elem mismatch. expected: %d got: %d", sorted[idx], toSort[idx])
			t.Fail()
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	toSort := make([][]int, b.N, b.N)
	for i := 0; i < b.N; i++ {
		toSort[i] = []int{31, 41, 59, 26, 53, 58, 97, 93, 23, 84, 2, 1, 112, 123, 12, 436, 23, 654, 654}
	}

	var sort []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bench1.InsertionSort(toSort[i])
		sort = toSort[i]
	}
	b.StopTimer()
	sorted = sort
}

func BenchmarkMySort(b *testing.B) {
	toSort := make([][]int, b.N, b.N)
	for i := 0; i < b.N; i++ {
		toSort[i] = []int{31, 41, 59, 26, 53, 58, 97, 93, 23, 84, 2, 1, 112, 123, 12, 436, 23, 654, 654}
	}

	var sort []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort = bench1.MySort(toSort[i])
	}
	b.StopTimer()
	sorted = sort
}

func BenchmarkMySort2(b *testing.B) {
	toSort := make([][]int, b.N, b.N)
	for i := 0; i < b.N; i++ {
		toSort[i] = []int{31, 41, 59, 26, 53, 58, 97, 93, 23, 84, 2, 1, 112, 123, 12, 436, 23, 654, 654}
	}

	var sort []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort = bench1.MySort2(toSort[i])
	}
	b.StopTimer()
	sorted = sort
}
