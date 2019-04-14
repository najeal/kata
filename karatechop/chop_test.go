package karatechop

import "testing"

func chopBenchmarkGeneric(b *testing.B, chopMethod ChopMethod) {
	array := []int{1, 3, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < b.N; i++ {
		chopMethod.Chop(20, array)
	}
}

func chopBenchmarkGenericLong(b *testing.B, chopMethod ChopMethod) {
	array := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		chopMethod.Chop(20, array)
	}
}

// ChopTestGeneric tests the Chop function for a given  ChopMethod
func chopTestGeneric(t *testing.T, chopMethod ChopMethod) {
	array := []int{}
	if -1 != chopMethod.Chop(3, array) {
		t.Errorf(`-1 != chop(3, %v)`, array)
	}

	array = []int{1}
	if -1 != chopMethod.Chop(3, array) {
		t.Errorf(`-1 != chop(3, %v)`, array)
	}
	if 0 != chopMethod.Chop(1, array) {
		t.Errorf(`0 != chop(1, %v)`, array)
	}

	array = []int{1, 3, 5}
	if 0 != chopMethod.Chop(1, array) {
		t.Errorf(`0 != chop(1, %v)`, array)
	}
	if 1 != chopMethod.Chop(3, array) {
		t.Errorf(`1 != chop(3, %v)`, array)
	}
	if 2 != chopMethod.Chop(5, array) {
		t.Errorf(`2 != chop(5, %v)`, array)
	}
	if -1 != chopMethod.Chop(0, array) {
		t.Errorf(`-1 != chop(0, %v)`, array)
	}
	if -1 != chopMethod.Chop(2, array) {
		t.Errorf(`-1 != chop(2, %v)`, array)
	}
	if -1 != chopMethod.Chop(4, array) {
		t.Errorf(`-1 != chop(4, %v)`, array)
	}
	if -1 != chopMethod.Chop(6, array) {
		t.Errorf(`-1 != chop(6, %v)`, array)
	}

	array = []int{1, 3, 5, 7}
	if 0 != chopMethod.Chop(1, array) {
		t.Errorf(`0 != chop(1, %v)`, array)
	}
	if 1 != chopMethod.Chop(3, array) {
		t.Errorf(`1 != chop(3, %v)`, array)
	}
	if 2 != chopMethod.Chop(5, array) {
		t.Errorf(`2 != chop(5, %v)`, array)
	}
	if 3 != chopMethod.Chop(7, array) {
		t.Errorf(`3 != chop(7, %v)`, array)
	}
	if -1 != chopMethod.Chop(0, array) {
		t.Errorf(`-1 != chop(0, %v)`, array)
	}
	if -1 != chopMethod.Chop(2, array) {
		t.Errorf(`-1 != chop(2, %v)`, array)
	}
	if -1 != chopMethod.Chop(4, array) {
		t.Errorf(`-1 != chop(4, %v)`, array)
	}
	if -1 != chopMethod.Chop(6, array) {
		t.Errorf(`-1 != chop(6, %v)`, array)
	}
	if -1 != chopMethod.Chop(8, array) {
		t.Errorf(`-1 != chop(8, %v)`, array)
	}
}
