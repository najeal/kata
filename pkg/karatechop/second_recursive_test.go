package karatechop

import "testing"

func TestSecondRecursiveChop(t *testing.T) {
	chopTestGeneric(t, new(SecondRecursiveChop))
}

func BenchmarkSecondRecursiveChop(b *testing.B) {
	chopBenchmarkGeneric(b, new(SecondRecursiveChop))
}

func BenchmarkSecondRecursiveChopLong(b *testing.B) {
	chopBenchmarkGenericLong(b, new(SecondRecursiveChop))
}
