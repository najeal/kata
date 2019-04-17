package anagrams

import (
	"fmt"

	"github.com/najeal/kata/pkg/common"
)

// NewAnagramDispatcher return a new AnagramDispatcher instance
func NewAnagramDispatcher(sortExtractor StringExtractorMethod) *AnagramDispatcher {
	return &AnagramDispatcher{
		wmap:          make(map[string][]string),
		sortExtractor: sortExtractor}
}

// NewSizeDispatcher return a new SizeDispatcher instance
func NewSizeDispatcher(sizeExtractor NumberExtractorMethod, wordCleander common.Cleaner) *SizeDispatcher {
	return &SizeDispatcher{
		smap:          make(map[int]*AnagramDispatcher),
		sizeExtractor: sizeExtractor,
		cleaner:       wordCleander}
}

// AnagramDispatcher is a string->[]string map
type AnagramDispatcher struct {
	wmap          map[string][]string
	sortExtractor StringExtractorMethod
}

// Add insert word in corresponding anagram set
func (an *AnagramDispatcher) Add(cleanedWord, word string) []string {
	key := an.sortExtractor.KeyFrom(cleanedWord)
	val := an.wmap[key]
	// if key does not exist, val will have default value empty slice
	newVal := append(val, word)
	an.wmap[key] = newVal
	return newVal
}

// PrintAll use printer to Print each anagram set
func (an *AnagramDispatcher) PrintAll(printer Printer) {
	for _, anagram := range an.wmap {
		printer.PrintAnagram(anagram)
	}
}

// SizeDispatcher is a int->Anagrams map
type SizeDispatcher struct {
	totalWords    int
	maxSize       int
	maxSet        []string
	smap          map[int]*AnagramDispatcher
	sizeExtractor NumberExtractorMethod
	cleaner       common.Cleaner
}

// Add insert word in corresponding AnagramDispatcher
func (sd *SizeDispatcher) Add(word string) {
	sd.totalWords++
	cleanedWord := sd.cleaner.Clean(word)
	size := len(cleanedWord)
	if size > sd.maxSize {
		sd.maxSize = size
	}
	anagramDispatcher := sd.smap[size]
	if anagramDispatcher == nil {
		anagramDispatcher = NewAnagramDispatcher(new(SortExtractor))
		sd.smap[size] = anagramDispatcher
	}
	anagramSet := anagramDispatcher.Add(cleanedWord, word)
	if len(anagramSet) > len(sd.maxSet) {
		sd.maxSet = anagramSet
	}
}

// PrintAll user printer to print each AnagramDispatcher
func (sd *SizeDispatcher) PrintAll(printer Printer) {
	printer.Print(" -> Printing all sets...")
	for _, value := range sd.smap {
		value.PrintAll(printer)
	}
}

// PrintLongestWordsAnagrams print longest anagram sets
func (sd *SizeDispatcher) PrintLongestWordsAnagrams(printer Printer) {
	printer.Print(" -> Printing all sets for longest words...")
	anagramDispatcher := sd.smap[sd.maxSize]
	anagramDispatcher.PrintAll(printer)
}

// PrintLongestSet prints set with longest word
func (sd *SizeDispatcher) PrintLongestSet(printer Printer) {
	printer.Print(" -> Printing longest set...")
	printer.PrintAnagram(sd.maxSet)
}

// PrintTotalWords prints total number of words stored
func (sd *SizeDispatcher) PrintTotalWords(printer Printer) {
	printer.Print(fmt.Sprintf(" -> Total words: %v", sd.totalWords))
}

// PrintTotalSets prints total number of sets stored
func (sd *SizeDispatcher) PrintTotalSets(printer Printer) {
	count := 0
	for _, elem := range sd.smap {
		for range elem.wmap {
			count++
		}
	}
	printer.Print(fmt.Sprintf(" -> Total sets: %v", count))
}
