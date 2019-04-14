package karatechop

import "testing"

func TestFunctionalChopChop(t *testing.T) {
	chopTestGeneric(t, new(FunctionalChop))
}

func BenchmarkFunctionalChopChop(b *testing.B) {
	chopBenchmarkGeneric(b, new(FunctionalChop))
}

func BenchmarkFunctionalChopChopLong(b *testing.B) {
	chopBenchmarkGenericLong(b, new(FunctionalChop))
}
