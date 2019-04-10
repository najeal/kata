package karatechop

import "testing"

func TestRecursiveChop(t *testing.T) {
	chopTestGeneric(t, new(RecursiveChop))
}
