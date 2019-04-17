package chain

import "fmt"

// NewSet creates new Set
func NewSet() *Set {
	return &Set{
		set: make(map[string]bool),
	}
}

// Set is a string->bool map
type Set struct {
	set map[string]bool
}

// Add insert new value in the Set
func (s *Set) Add(val string) {
	s.set[val] = true
}

// IsIn returns if value is present or not in the Set
func (s *Set) IsIn(val string) (res bool) {
	_, res = s.set[val]
	return
}

// GetValues returns values stored in the Set
func (s *Set) GetValues() []string {
	values := make([]string, len(s.set))
	index := 0
	for key := range s.set {
		values[index] = key
		index++
	}
	return values
}

// String format values in string
func (s *Set) String() string {
	return fmt.Sprintf("%v", s.set)
}
