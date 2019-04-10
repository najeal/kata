package karatechop

import "testing"

func TestIterativeChop(t *testing.T) {
	chopTestGeneric(t, new(IterativeChop))
}
