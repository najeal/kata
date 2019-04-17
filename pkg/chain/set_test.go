package chain

import (
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet()
	if set == nil {
		t.Errorf("Set created by NewSet function should not be nil, but it is")
	}
	set.Add("work")
	set.Add("work")

	if true == set.IsIn("newway") {
		t.Errorf("Set IsIn newway should return false")
	}

	if false == set.IsIn("work") {
		t.Errorf("Set IsIn work should return true")
	}

	set.Add("newway")

	if false == set.IsIn("newway") {
		t.Errorf("Set IsIn newway should return true")
	}

	values := set.GetValues()
	length := len(values)
	if length != 2 {
		t.Errorf("Set GetValues should return only 2 values but returned %v", length)
	} else {
		if false == (values[0] == "work" && values[1] == "newway" || values[0] == "newway" && values[1] == "work") {
			t.Errorf("values returned by Set GetValues are incorrects: %v", values)
		}
	}
}
