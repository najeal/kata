package karatechop

import "testing"

func TestIterativeChop(t *testing.T) {
	chopTestGeneric(t, new(IterativeChop))
}

func BenchmarkIterativeChop(b *testing.B) {
	chopBenchmarkGeneric(b, new(IterativeChop))
}

func BenchmarkIterativeChopLong(b *testing.B) {
	chopBenchmarkGenericLong(b, new(IterativeChop))
}
