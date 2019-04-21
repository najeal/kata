package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	store, err := NewStore("golang", "langgo")
	assert.Nil(t, err)

	store.Add("restau") // distance 6
	assert.Equal(t, 1, len(store.smap))

	store.Add("geltau") // distance 4
	assert.Equal(t, 2, len(store.smap))

	store.Add("rowaqg") // distance 3
	assert.Equal(t, 3, len(store.smap))

	store.Add("rowaqa") // distance 4
	assert.Equal(t, 3, len(store.smap))

	assert.Equal(t, 0, len(store.smap[1]))
	assert.Equal(t, 2, len(store.smap[4]))
	assert.Equal(t, 6, store.maxDistance)
}

func TestStoreFeedGraph(t *testing.T) {
	store, _ := NewStore("cat", "dog")
	store.Add("cot")
	store.Add("cog")
	store.Add("cos")
	store.Add("rps")
	store.Add("cps")
	graph := &MatrixGraph{matrix: make(map[string]*Set)}
	store.FeedGraph(graph)
	res, _ := graph.BFS("cat", "dog")
	expectedRes := []string{"cat", "cot", "cog", "dog"}
	assert.Equal(t, expectedRes, res)
}
