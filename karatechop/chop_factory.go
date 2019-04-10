package karatechop

import (
	"fmt"
)

const (
	// IterativeChopType used to identify IterativeChop method
	IterativeChopType = 1
	// RecursiveChopType used to identify RecusiveChop method
	RecursiveChopType = 2
)

// ChopMethod is used to find an the index of a number in a sorted array source
type ChopMethod interface {
	Chop(target int, source []int) int
}

// GetChopMethod generates the specific ChopMethod asked
func GetChopMethod(method int) (ChopMethod, error) {
	switch method {
	case IterativeChopType:
		return new(IterativeChop), nil
	case RecursiveChopType:
		return new(RecursiveChop), nil
	default:
		return nil, fmt.Errorf("This ChopMethod is not supported")
	}
}
