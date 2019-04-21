package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphNew(t *testing.T) {
	graph := NewMatrixGraph()
	assert.NotNil(t, graph)
	assert.NotNil(t, graph.matrix)
}
func TestGraphAddUniEdge(t *testing.T) {
	graph := NewMatrixGraph()
	graph.AddUniDirectionEdge("a", "a")
	assert.Equal(t, 1, len(graph.matrix))

	graph.AddUniDirectionEdge("a", "b")
	assert.Equal(t, 1, len(graph.matrix))

	graph.AddUniDirectionEdge("b", "c")
	assert.Equal(t, 2, len(graph.matrix))
}

func TestGraphAddMultiEdgeAndBFS(t *testing.T) {
	graph := NewMatrixGraph()
	graph.AddMultiDirectionEdge("a", "a")
	assert.Equal(t, 1, len(graph.matrix))

	graph.AddMultiDirectionEdge("c", "d")
	assert.Equal(t, 3, len(graph.matrix))

	graph.AddMultiDirectionEdge("a", "d")
	assert.Equal(t, 3, len(graph.matrix))

	res, err := graph.BFS("a", "c")
	assert.Nil(t, err)

	expectedRes := []string{"a", "d", "c"}
	assert.Equal(t, expectedRes, res)
}
