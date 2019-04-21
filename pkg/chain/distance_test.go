package chain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordDistancer(t *testing.T) {
	tests := []struct {
		name         string
		wordA        string
		wordB        string
		expectedErr  error
		expectedDist int
	}{
		{
			"not same length",
			"a",
			"",
			fmt.Errorf("Cannot search distance for different word sizes"),
			0,
		},
		{
			"same empty",
			"",
			"",
			nil,
			0,
		},
		{
			"one letter length dist 1",
			"a",
			"c",
			nil,
			1,
		},
		{
			"four letter length dist 1",
			"abct",
			"cbct",
			nil,
			1,
		},
		{
			"four letter length dist 3",
			"bcat",
			"cbct",
			nil,
			3,
		},
	}

	for _, utest := range tests {
		t.Run(utest.name, func(t *testing.T) {
			distancer := new(WordDistancer)
			dist, err := distancer.findDistance(utest.wordA, utest.wordB)
			assert.Equal(t, utest.expectedErr, err)
			assert.Equal(t, utest.expectedDist, dist)
		})
	}
}
