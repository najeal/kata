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

	_, err = GetChopMethod(FunctionalChopType)
	if err != nil {
		t.Error("A FunctionalChopType should exist")
	}

	_, err = GetChopMethod(SecondRecursiveChopType)
	if err != nil {
		t.Error("A SecondRecursiveChopType should exist")
	}

	_, err = GetChopMethod(SecondFunctionalChopType)
	if err != nil {
		t.Error("A SecondFunctionalChopType should exist")
	}
}
