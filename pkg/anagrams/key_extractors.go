package anagrams

// StringExtractorMethod is used to extract key string from string
type StringExtractorMethod interface {
	KeyFrom(a string) string
}

// SortExtractor is used to extract key from string
type SortExtractor struct{}

// KeyFrom extract ordered word letters from input word
func (se *SortExtractor) KeyFrom(word string) string {
	newWord := []rune(word)
	return string(se.sortFusion(newWord))
}

func (se *SortExtractor) sortFusion(word []rune) []rune {
	maxIndex := len(word) - 1
	if maxIndex < 1 {
		return word
	}
	return se.fusion(se.sortFusion(word[0:maxIndex/2+1]), se.sortFusion(word[maxIndex/2+1:maxIndex+1]))
}

func (se *SortExtractor) fusion(leftWord, rightWord []rune) []rune {
	if len(leftWord) < 1 {
		return rightWord
	} else if len(rightWord) < 1 {
		return leftWord
	} else if leftWord[0] < rightWord[0] {
		return append([]rune{leftWord[0]}, se.fusion(leftWord[1:], rightWord)...)
	} else {
		return append([]rune{rightWord[0]}, se.fusion(leftWord, rightWord[1:])...)
	}
}

// NumberExtractorMethod is used to extract key number from string
type NumberExtractorMethod interface {
	KeyFrom(a string) int
}

// LenExtractor is used to extract length from string
type LenExtractor struct{}

// KeyFrom extract word length
func (le *LenExtractor) KeyFrom(word string) int {
	return len(word)
}
