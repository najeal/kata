package karatechop

import "testing"

func TestSecondFunctionalChop(t *testing.T) {
	chopTestGeneric(t, new(SecondFunctionalChop))
}

func BenchmarkSecondFunctionalChop(b *testing.B) {
	chopBenchmarkGeneric(b, new(SecondFunctionalChop))
}

func BenchmarkSecondFunctionalChopLong(b *testing.B) {
	chopBenchmarkGenericLong(b, new(SecondFunctionalChop))
}
