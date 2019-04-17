package chain

import (
	"fmt"

	"github.com/najeal/kata/pkg/common"
)

// NewStore returns new Store instance
func NewStore(start, end string) (*Store, error) {
	if len(start) != len(end) {
		return nil, fmt.Errorf("start and end word should have the same length")
	}
	distancer := NewWordDistancer()
	cleaner := common.NewWordCleaner()
	return &Store{
		start:     start,
		end:       end,
		length:    len(start),
		cleaner:   cleaner,
		distancer: distancer,
		smap:      make(map[int][]string),
	}, nil
}

// Store  is a int->[]string map
type Store struct {
	start       string
	end         string
	length      int
	maxDistance int
	cleaner     common.Cleaner
	distancer   Distancer
	smap        map[int][]string
}

// Add inserts word in the store
func (s *Store) Add(word string) {
	cWord := s.cleaner.Clean(word)
	if len(word) != s.length {
		return
	}
	distance, err := s.distancer.findDistance(s.start, cWord)
	if err != nil {
		return
	}
	s.saveGreater(distance)
	s.insertMap(distance, word)
	return
}

func (s *Store) saveGreater(distance int) {
	if distance > s.maxDistance {
		s.maxDistance = distance
	}
}

func (s *Store) insertMap(index int, word string) {
	words := s.smap[index]
	s.smap[index] = append(words, word)
}

// FeedGraph adds nodes and edges to graph
func (s *Store) FeedGraph(graph Graph) {
	s.feedGraphFromStart(graph)
	for i := 1; i <= s.maxDistance; i++ {
		parents := s.smap[i]
		potentialChildren := parents
		s.feedGraph(graph, parents, potentialChildren)
		if i < s.maxDistance {
			potentialChildren = s.smap[i+1]
			s.feedGraph(graph, parents, potentialChildren)
		}
	}
}

func (s *Store) feedGraphFromStart(graph Graph) {
	if s.maxDistance > 1 {
		for _, firstChildren := range s.smap[1] {
			graph.AddUniDirectionEdge(s.start, firstChildren)
		}
	}
}

func (s *Store) feedGraph(graph Graph, parents, potentialChildren []string) {
	for _, parentSource := range parents {
		if dist, err := s.distancer.findDistance(parentSource, s.end); err == nil && dist == 1 {
			graph.AddUniDirectionEdge(parentSource, s.end)
		}
		// feed parents with parents
		for _, parent := range parents {
			if dist, err := s.distancer.findDistance(parentSource, parent); err == nil && dist == 1 {
				graph.AddMultiDirectionEdge(parentSource, parent)
			}
		}
		// feed parents with child
		for _, child := range potentialChildren {
			if dist, err := s.distancer.findDistance(parentSource, child); err == nil && dist == 1 {
				graph.AddMultiDirectionEdge(parentSource, child)
			}
		}
	}
}
