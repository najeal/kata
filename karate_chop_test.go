package kata

import "testing"

// TestChop tests the chop function
func TestChop(t *testing.T) {
	array := []int{}
	if -1 != chop(3, array) {
		t.Errorf(`-1 != chop(3, %v)`, array)
	}

	array = []int{1}
	if -1 != chop(3, array) {
		t.Errorf(`-1 != chop(3, %v)`, array)
	}
	if 0 != chop(1, array) {
		t.Errorf(`0 != chop(1, %v)`, array)
	}

	array = []int{1, 3, 5}
	if 0 != chop(1, array) {
		t.Errorf(`0 != chop(1, %v)`, array)
	}
	if 1 != chop(3, array) {
		t.Errorf(`1 != chop(3, %v)`, array)
	}
	if 2 != chop(5, array) {
		t.Errorf(`2 != chop(5, %v)`, array)
	}
	if -1 != chop(0, array) {
		t.Errorf(`-1 != chop(0, %v)`, array)
	}
	if -1 != chop(2, array) {
		t.Errorf(`-1 != chop(2, %v)`, array)
	}
	if -1 != chop(4, array) {
		t.Errorf(`-1 != chop(4, %v)`, array)
	}
	if -1 != chop(6, array) {
		t.Errorf(`-1 != chop(6, %v)`, array)
	}

	array = []int{1, 3, 5, 7}
	if 0 != chop(1, array) {
		t.Errorf(`0 != chop(1, %v)`, array)
	}
	if 1 != chop(3, array) {
		t.Errorf(`1 != chop(3, %v)`, array)
	}
	if 2 != chop(5, array) {
		t.Errorf(`2 != chop(5, %v)`, array)
	}
	if 3 != chop(7, array) {
		t.Errorf(`3 != chop(7, %v)`, array)
	}
	if -1 != chop(0, array) {
		t.Errorf(`-1 != chop(0, %v)`, array)
	}
	if -1 != chop(2, array) {
		t.Errorf(`-1 != chop(2, %v)`, array)
	}
	if -1 != chop(4, array) {
		t.Errorf(`-1 != chop(4, %v)`, array)
	}
	if -1 != chop(6, array) {
		t.Errorf(`-1 != chop(6, %v)`, array)
	}
	if -1 != chop(8, array) {
		t.Errorf(`-1 != chop(8, %v)`, array)
	}
}
