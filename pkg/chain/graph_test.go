package chain

import (
	"reflect"
	"testing"
)

func TestGraphNew(t *testing.T) {
	graph := NewMatrixGraph()
	if graph == nil {
		t.Errorf("graph should not be nil")
	}
	if graph.matrix == nil {
		t.Errorf("matrix graph should not be nil")
	}
}
func TestGraphAddUniEdge(t *testing.T) {
	graph := NewMatrixGraph()
	graph.AddUniDirectionEdge("a", "a")
	if length := len(graph.matrix); length != 1 {
		t.Errorf("graph matrix length should be 1, but is %v", length)
	}

	graph.AddUniDirectionEdge("a", "b")
	if length := len(graph.matrix); length != 1 {
		t.Errorf("graph matrix length should be 1, but is %v", length)
	}

	graph.AddUniDirectionEdge("b", "c")
	if length := len(graph.matrix); length != 2 {
		t.Errorf("graph matrix length should be 2, but is %v", length)
	}
}

func TestGraphAddMultiEdge(t *testing.T) {
	graph := NewMatrixGraph()
	graph.AddMultiDirectionEdge("a", "a")
	if length := len(graph.matrix); length != 1 {
		t.Errorf("graph matrix length should be 1, but is %v", length)
	}

	graph.AddMultiDirectionEdge("c", "d")
	if length := len(graph.matrix); length != 3 {
		t.Errorf("graph matrix length should be 3, but is %v", length)
	}

	graph.AddMultiDirectionEdge("a", "d")
	if length := len(graph.matrix); length != 3 {
		t.Errorf("graph matrix length should be 3, but is %v", length)
	}

	res, err := graph.BFS("a", "c")
	if err != nil {
		t.Errorf("graph bfs failed, but it should not")
	}

	expectedRes := []string{"a", "d", "c"}
	if reflect.DeepEqual(expectedRes, res) == false {
		t.Errorf("graph bfs result is %v but should be %v", res, expectedRes)
	}
}
