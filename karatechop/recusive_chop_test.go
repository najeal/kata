package karatechop

import "testing"

func TestRecursiveChop(t *testing.T) {
	chopTestGeneric(t, new(RecursiveChop))
}

func BenchmarkRecursiveChop(b *testing.B) {
	chopBenchmarkGeneric(b, new(RecursiveChop))
}

func BenchmarkRecursiveChopLong(b *testing.B) {
	chopBenchmarkGenericLong(b, new(RecursiveChop))
}
