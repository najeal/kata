package anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/najeal/kata/pkg/common"
)

// TestAnagramDispatcher tests Add function of AnagramDispatcher
func TestAnagramDispatcher(t *testing.T) {
	anDispatcher := NewAnagramDispatcher(new(SortExtractor))
	set := anDispatcher.Add("velo", "velo")
	assert.Equal(t, 1, len(set))
	assert.Equal(t, "velo", set[0])

	set = anDispatcher.Add("love", "love")
	assert.Equal(t, 2, len(set))
	assert.Equal(t, "velo", set[0])
	assert.Equal(t, "love", set[1])

	assert.Equal(t, 1, len(anDispatcher.wmap))
}

// TestSizeDispatcher tests Add function of SizeDispatcher
func TestSizeDispatcher(t *testing.T) {
	sizeDispatcher := NewSizeDispatcher(new(LenExtractor), new(common.WordCleaner))
	assert.Equal(t, 0, len(sizeDispatcher.maxSet))
	assert.Equal(t, 0, sizeDispatcher.maxSize)

	sizeDispatcher.Add("velo")
	assert.Equal(t, 1, len(sizeDispatcher.smap))

	sizeDispatcher.Add("love")
	assert.Equal(t, 1, len(sizeDispatcher.smap))

	sizeDispatcher.Add("lover")
	assert.Equal(t, 2, len(sizeDispatcher.smap))

	sizeDispatcher.Add("the Morse code")
	sizeDispatcher.Add("Here come dots")
	assert.Equal(t, 3, len(sizeDispatcher.smap))
}

// TestSizeDispatcherLongestAndMaxSize tests that longest anagram set and max word size are nicely stored
func TestSizeDispatcherLongestAndMaxSize(t *testing.T) {
	sizeDispatcher := NewSizeDispatcher(new(LenExtractor), new(common.WordCleaner))
	sizeDispatcher.Add("velo")
	sizeDispatcher.Add("love")
	sizeDispatcher.Add("lover")
	expectedLongest := []string{"velo", "love"}
	assert.Equal(t, expectedLongest, sizeDispatcher.maxSet)

	sizeDispatcher.Add("rlove")
	assert.Equal(t, expectedLongest, sizeDispatcher.maxSet)

	sizeDispatcher.Add("lrove")
	expectedLongest = []string{"lover", "rlove", "lrove"}
	assert.Equal(t, expectedLongest, sizeDispatcher.maxSet)

	assert.Equal(t, 5, sizeDispatcher.maxSize)

	sizeDispatcher.Add("unpolitic")
	assert.Equal(t, 9, sizeDispatcher.maxSize)

	expectedLongest = []string{"lover", "rlove", "lrove"}
	assert.Equal(t, expectedLongest, sizeDispatcher.maxSet)

	sizeDispatcher.Add("the Morse code")
	assert.Equal(t, 12, sizeDispatcher.maxSize)
}
