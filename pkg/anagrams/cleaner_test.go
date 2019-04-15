package anagrams

import (
	"testing"
)

func TestWordCleaner(t *testing.T) {
	wordToClean := "@te[as'asd{"
	expectedWord := "teasasd"
	wordCleaner := new(WordCleaner)
	wordResult := wordCleaner.Clean(wordToClean)
	if wordResult != expectedWord {
		t.Errorf("Cleaned word %v is different from %v expected", wordResult, expectedWord)
	}

	wordToClean = "interlink's"
	expectedWord = "interlinks"
	wordResult = wordCleaner.Clean(wordToClean)
	if wordResult != expectedWord {
		t.Errorf("Cleaned word %v is different from %v expected", wordResult, expectedWord)
	}

	wordToClean = "the Morse code"
	expectedWord = "themorsecode"
	wordResult = wordCleaner.Clean(wordToClean)
	if wordResult != expectedWord {
		t.Errorf("Cleaned word %v is different from %v expected", wordResult, expectedWord)
	}

	wordToClean = "Here come dots"
	expectedWord = "herecomedots"
	wordResult = wordCleaner.Clean(wordToClean)
	if wordResult != expectedWord {
		t.Errorf("Cleaned word %v is different from %v expected", wordResult, expectedWord)
	}
}
