package karatechop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func chopBenchmarkGeneric(b *testing.B, chopMethod ChopMethod) {
	array := []int{1, 3, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < b.N; i++ {
		chopMethod.Chop(20, array)
	}
}

func chopBenchmarkGenericLong(b *testing.B, chopMethod ChopMethod) {
	array := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		chopMethod.Chop(20, array)
	}
}

// ChopTestGeneric tests the Chop function for a given  ChopMethod
func chopTestGeneric(t *testing.T, chopMethod ChopMethod) {
	tests := []struct {
		name           string
		array          []int
		search         []int
		resultExpected []int
	}{
		{
			"empty array",
			[]int{},
			[]int{3},
			[]int{-1},
		},
		{
			"one element array",
			[]int{1},
			[]int{3, 1},
			[]int{-1, 0},
		},
		{
			"three elements array",
			[]int{1, 3, 5},
			[]int{1, 3, 5, 0, 2, 4, 6},
			[]int{0, 1, 2, -1, -1, -1, -1},
		},
		{
			"four elements array",
			[]int{1, 3, 5, 7},
			[]int{1, 3, 5, 7, 0, 2, 4, 6, 8},
			[]int{0, 1, 2, 3, -1, -1, -1, -1, -1},
		},
	}

	for _, utest := range tests {
		t.Run(utest.name, func(t *testing.T) {
			for index := range utest.search {
				res := chopMethod.Chop(utest.search[index], utest.array)
				assert.Equal(t, utest.resultExpected[index], res)
			}
		})
	}
}
