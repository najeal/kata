package anagrams

import (
	"reflect"
	"strings"
	"testing"

	"github.com/najeal/kata/pkg/common"
)

// TestAnagramDispatcher tests Add function of AnagramDispatcher
func TestAnagramDispatcher(t *testing.T) {
	anDispatcher := NewAnagramDispatcher(new(SortExtractor))
	set := anDispatcher.Add("velo", "velo")
	if len(set) != 1 {
		t.Errorf("first call anagram add: set length should be 1")
	} else {
		if strings.Compare(set[0], "velo") != 0 {
			t.Errorf("first call anagram add: first elem should be velo but is %v", set[0])
		}
	}
	set = anDispatcher.Add("love", "love")
	if len(set) != 2 {
		t.Errorf("second call anagram add: set length should be 2")
	} else {
		if strings.Compare(set[0], "velo") != 0 {
			t.Errorf("second call anagram add: first elem should be velo but is %v", set[0])
		}
		if strings.Compare(set[1], "love") != 0 {
			t.Errorf("second call anagram add: second elem should be love but is %v", set[0])
		}
	}
	anagramsCount := len(anDispatcher.wmap)
	if anagramsCount != 1 {
		t.Errorf("anagram count should be 1 but is %v", anagramsCount)
	}
}

// TestSizeDispatcher tests Add function of SizeDispatcher
func TestSizeDispatcher(t *testing.T) {
	sizeDispatcher := NewSizeDispatcher(new(LenExtractor), new(common.WordCleaner))
	if len(sizeDispatcher.maxSet) > 0 {
		t.Errorf("longest set length should be 0")
	}
	if sizeDispatcher.maxSize > 0 {
		t.Errorf("longest word should be 0")
	}
	sizeDispatcher.Add("velo")
	differentSizesCount := len(sizeDispatcher.smap)
	if differentSizesCount != 1 {
		t.Errorf("size count should be 1 but is %v", differentSizesCount)
	}
	sizeDispatcher.Add("love")
	differentSizesCount = len(sizeDispatcher.smap)
	if differentSizesCount != 1 {
		t.Errorf("size count should be 1 but is %v", differentSizesCount)
	}
	sizeDispatcher.Add("lover")
	differentSizesCount = len(sizeDispatcher.smap)
	if differentSizesCount != 2 {
		t.Errorf("size count should be 2 but is %v", differentSizesCount)
	}

	sizeDispatcher.Add("the Morse code")
	sizeDispatcher.Add("Here come dots")
	differentSizesCount = len(sizeDispatcher.smap)
	if differentSizesCount != 3 {
		t.Errorf("size count should be 3 but is %v", differentSizesCount)
	}
}

// TestSizeDispatcherLongestAndMaxSize tests that longest anagram set and max word size are nicely stored
func TestSizeDispatcherLongestAndMaxSize(t *testing.T) {
	sizeDispatcher := NewSizeDispatcher(new(LenExtractor), new(common.WordCleaner))
	sizeDispatcher.Add("velo")
	sizeDispatcher.Add("love")
	sizeDispatcher.Add("lover")
	expectedLongest := []string{"velo", "love"}
	if reflect.DeepEqual(sizeDispatcher.maxSet, expectedLongest) != true {
		t.Errorf("longest is %v whereas we should have %v", sizeDispatcher.maxSet, expectedLongest)
	}
	sizeDispatcher.Add("rlove")
	if reflect.DeepEqual(sizeDispatcher.maxSet, expectedLongest) != true {
		t.Errorf("longest is %v whereas we should have %v", sizeDispatcher.maxSet, expectedLongest)
	}
	sizeDispatcher.Add("lrove")
	expectedLongest = []string{"lover", "rlove", "lrove"}
	if reflect.DeepEqual(sizeDispatcher.maxSet, expectedLongest) != true {
		t.Errorf("longest is %v whereas we should have %v", sizeDispatcher.maxSet, expectedLongest)
	}
	maxSize := sizeDispatcher.maxSize
	expectedMaxSize := 5
	if maxSize != expectedMaxSize {
		t.Errorf("longest word length is %v but should be %v", maxSize, expectedMaxSize)
	}
	sizeDispatcher.Add("unpolitic")
	maxSize = sizeDispatcher.maxSize
	expectedMaxSize = 9
	if maxSize != expectedMaxSize {
		t.Errorf("longest word length is %v but should be %v", maxSize, expectedMaxSize)
	}

	expectedLongest = []string{"lover", "rlove", "lrove"}
	if reflect.DeepEqual(sizeDispatcher.maxSet, expectedLongest) != true {
		t.Errorf("longest is %v whereas we should have %v", sizeDispatcher.maxSet, expectedLongest)
	}

	sizeDispatcher.Add("the Morse code")
	maxSize = sizeDispatcher.maxSize
	expectedMaxSize = 12
	if maxSize != expectedMaxSize {
		t.Errorf("longest word length is %v but should be %v", maxSize, expectedMaxSize)
	}
}
