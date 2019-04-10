package karatechop

import (
	"testing"
)

func TestChopFactory(t *testing.T) {
	_, err := GetChopMethod(RecursiveChopType)
	if err != nil {
		t.Error("A RecursiveChopType should exist")
	}

	_, err = GetChopMethod(IterativeChopType)
	if err != nil {
		t.Error("A IterativeChopType should exist")
	}
}
