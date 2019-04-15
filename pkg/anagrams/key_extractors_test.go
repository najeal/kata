package anagrams

import (
	"strings"
	"testing"
)

func TestSortWord(t *testing.T) {
	sExtractor := new(SortExtractor)
	errorMessage := "sorted word %v got is different to %v"
	word := "cowboy"
	expectedWord := "bcoowy"
	res := sExtractor.KeyFrom(word)
	if 0 != strings.Compare(res, expectedWord) {
		t.Errorf(errorMessage, res, expectedWord)
	}

	word = "c"
	expectedWord = "c"
	res = sExtractor.KeyFrom(word)
	if 0 != strings.Compare(res, expectedWord) {
		t.Errorf(errorMessage, res, expectedWord)
	}

	word = "ca"
	expectedWord = "ac"
	res = sExtractor.KeyFrom(word)
	if 0 != strings.Compare(res, expectedWord) {
		t.Errorf(errorMessage, res, expectedWord)
	}

	word = "themorsecode"
	expectedWord = "cdeeehmoorst"
	res = sExtractor.KeyFrom(word)
	if 0 != strings.Compare(res, expectedWord) {
		t.Errorf(errorMessage, res, expectedWord)
	}

	word = "herecomedots"
	expectedWord = "cdeeehmoorst"
	res = sExtractor.KeyFrom(word)
	if 0 != strings.Compare(res, expectedWord) {
		t.Errorf(errorMessage, res, expectedWord)
	}
}
