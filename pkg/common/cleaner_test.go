package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCleaner(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			"delete special keys",
			"@te[as'asd{",
			"teasasd",
		},
		{
			"delete spaces",
			"the Morse code",
			"themorsecode",
		},
		{
			"upper to lower cases",
			"ABCD",
			"abcd",
		},
	}

	for _, utest := range tests {
		t.Run(utest.name, func(t *testing.T) {
			wordCleaner := new(WordCleaner)
			wordResult := wordCleaner.Clean(utest.input)
			assert.Equal(t, utest.expectedOutput, wordResult)
		})
	}
}
