package uset

import "testing"

// len, add, delete, find

func Test(t *testing.T) {
	set := make(USet)
	nVals := 5

	// test Add
	for i := 0; i < nVals; i++ {
		set.Add(i)
	}

	// test Size
	wantSize := nVals
	if got := set.Size(); got != wantSize {
		t.Errorf("set.Size() = %d want: %d\n", got, wantSize)
	}

	// re-add existing values
	for i := 0; i < nVals; i++ {
		set.Add(i)
	}
	wantSize = nVals
	if got := set.Size(); got != wantSize {
		t.Errorf("After adding existing values\n set.Size() = %d want: %d\n",
			got, wantSize)
	}

	// test Find
	for i := 0; i < nVals; i++ {
		x, found := set.Find(i)
		if found != true || x != i {
			t.Errorf("set.Find() failed to find value: %d got: %d\n", i, x)
		}
	}

	for i := nVals + 1; i < nVals*2; i++ {
		_, found := set.Find(i)
		if found != false {
			t.Errorf("set.Find() failed to find value: %d\n", i)
		}
	}

}
