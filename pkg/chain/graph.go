package chain

import (
	"container/list"
	"fmt"
)

// Graph represents Graph data structure
type Graph interface {
	AddUniDirectionEdge(parent, child string)
	AddMultiDirectionEdge(parent, child string)
	BFS(source, target string) ([]string, error)
}

// NewMatrixGraph creates new MatrixGraph
func NewMatrixGraph() *MatrixGraph {
	return &MatrixGraph{
		matrix: make(map[string]*Set),
	}
}

// MatrixGraph implements Graph interface with adjacency matrix method
type MatrixGraph struct {
	// matrix is an adjacency matrix
	matrix map[string]*Set
}

// AddUniDirectionEdge add adjacent nodes with uni directions
func (g *MatrixGraph) AddUniDirectionEdge(parent, child string) {
	children, ok := g.matrix[parent]
	if ok {
		children.Add(child)
	} else {
		newSet := NewSet()
		newSet.Add(child)
		g.matrix[parent] = newSet
	}
}

// AddMultiDirectionEdge add adjacent nodes with multi directions
func (g *MatrixGraph) AddMultiDirectionEdge(parent, child string) {
	g.AddUniDirectionEdge(parent, child)
	g.AddUniDirectionEdge(child, parent)
}

func (g *MatrixGraph) getChildrenFromParent(parent string) ([]string, bool) {
	children, ok := g.matrix[parent]
	if ok {
		return children.GetValues(), true
	}
	return []string{}, ok
}

// BFS searches shortest way in the graph from source to target
func (g *MatrixGraph) BFS(source string, target string) ([]string, error) {
	res := g.bfs(source, target)
	iterate := true
	chain := make([]string, 0)
	parent := target
	chain = append(chain, target)
	for iterate {
		child, ok := res[parent]
		if ok {
			chain = append(chain, child)
			parent = child
		} else {
			iterate = false
		}
	}
	if len(chain) < 2 {
		return nil, fmt.Errorf("Did not find chain")
	}
	// reverse chain
	for i := len(chain)/2 - 1; i >= 0; i-- {
		opp := len(chain) - 1 - i
		chain[i], chain[opp] = chain[opp], chain[i]
	}
	return chain, nil
}

func (g *MatrixGraph) bfs(source string, target string) map[string]string {
	parent := make(map[string]string)
	visited := make(map[string]bool)
	visited[source] = true
	queue := list.New()
	queue.PushBack(source)
	for current := queue.Front(); current != nil; current = queue.Front() {
		queue.Remove(current)
		children, _ := g.getChildrenFromParent(current.Value.(string))
		for _, child := range children {
			if _, ok := visited[child]; ok != true {
				visited[child] = true
				parent[child] = current.Value.(string)
				queue.PushBack(child)
				if child == target {
					return parent
				}
			}
		}
	}
	return parent
}
