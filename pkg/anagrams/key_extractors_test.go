package anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortWord(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		expected string
	}{
		{
			"normal word",
			"cowboy",
			"bcoowy",
		},
		{
			"long word 1",
			"themorsecode",
			"cdeeehmoorst",
		},
		{
			"long word 2",
			"herecomedots",
			"cdeeehmoorst",
		},
		{
			"one letter",
			"c",
			"c",
		},
		{
			"two letters",
			"ca",
			"ac",
		},
	}

	for _, utest := range tests {
		t.Run(utest.name, func(t *testing.T) {
			sExtractor := new(SortExtractor)
			res := sExtractor.KeyFrom(utest.word)
			assert.Equal(t, utest.expected, res)
		})
	}
}
