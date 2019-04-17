package chain

import (
	"log"
	"reflect"
	"testing"
)

func TestStoreInsert(t *testing.T) {
	store, err := NewStore("golang", "langgo")
	if err != nil {
		log.Println(err)
		t.Errorf("First Store creation should not return error")
	}

	store.Add("restau") // distance 6
	length := len(store.smap)
	if length != 1 {
		t.Errorf("smap length should be 1 but is %v", length)
	}

	store.Add("geltau") // distance 4
	length = len(store.smap)
	if length != 2 {
		t.Errorf("smap length should be 2 but is %v", length)
	}

	store.Add("rowaqg") // distance 3
	length = len(store.smap)
	if length != 3 {
		t.Errorf("smap length should be 3 but is %v", length)
	}

	store.Add("rowaqa") // distance 4
	length = len(store.smap)
	if length != 3 {
		t.Errorf("smap length should be 3 but is %v", length)
	}

	elems := store.smap[1]
	if numStored := len(elems); numStored != 0 {
		t.Errorf("numStored should be 0 but is %v", numStored)
	}

	elems = store.smap[4]
	if numStored := len(elems); numStored != 2 {
		t.Errorf("numStored should be 2 but is %v", numStored)
	}

	if maxDist := store.maxDistance; maxDist != 6 {
		t.Errorf("maxDist should be 6 but is %v", maxDist)
	}
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
	if reflect.DeepEqual(expectedRes, res) != true {
		t.Errorf("BFS search should give %v after the store has feeded the graph, but got", res)
	}
}
