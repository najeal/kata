package chain

import (
	"testing"
)

func TestWordDistancer(t *testing.T) {
	distancer := new(WordDistancer)

	dist, err := distancer.findDistance("a", "")
	if err == nil {
		t.Errorf("different word length should return error")
	}

	dist, err = distancer.findDistance("", "")
	if err != nil {
		t.Errorf("empties word should not return error")
	}
	if dist != 0 {
		t.Errorf("dist is %v but should be 0", dist)
	}

	dist, err = distancer.findDistance("a", "c")
	if err != nil {
		t.Errorf("empties word should not return error")
	}
	if dist != 1 {
		t.Errorf("dist is %v but should be 1", dist)
	}

	dist, err = distancer.findDistance("abct", "cbct")
	if err != nil {
		t.Errorf("empties word should not return error")
	}
	if dist != 1 {
		t.Errorf("dist is %v but should be 1", dist)
	}

	dist, err = distancer.findDistance("bcat", "cbct")
	if err != nil {
		t.Errorf("empties word should not return error")
	}
	if dist != 3 {
		t.Errorf("dist is %v but should be 3", dist)
	}

}
