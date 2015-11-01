package bench5_test

import (
	"testing"

	"github.com/dlsniper/codemotion/code/bench/5"
)

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench5.Format()
	}
}

func BenchmarkFormatNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench5.Format()
	}
}
