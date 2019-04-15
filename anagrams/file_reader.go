package anagrams

import (
	"bufio"
	"os"
)

// FileReader reads files
type FileReader struct{}

// FeedFromFile reads file and feeds each line
func (fr *FileReader) FeedFromFile(file *os.File, feed func(string)) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		feed(scanner.Text())
	}
}
