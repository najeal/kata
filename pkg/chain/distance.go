package chain

import "fmt"

// Distancer is used to find distance between two strings
type Distancer interface {
	findDistance(a, b string) (int, error)
}

// NewWordDistancer creates new WordDistancer
func NewWordDistancer() *WordDistancer {
	return &WordDistancer{}
}

// WordDistancer is used to find distance between two words
type WordDistancer struct{}

// findDistance search distance between lowercase words
func (wd *WordDistancer) findDistance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Cannot search distance for different word sizes")
	}
	ra, rb := []rune(a), []rune(b)
	distance := 0
	for index := range ra {
		if ra[index] != rb[index] {
			distance++
		}
	}
	return distance, nil
}
