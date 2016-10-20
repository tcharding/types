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

	// test Len
	wantLength := nVals
	if got := set.Len(); got != wantLength {
		t.Errorf("set.Len() = %d want: %d\n", got, wantLength)
	}

	// re-add existing values
	for i := 0; i < nVals; i++ {
		set.Add(i)
	}
	wantLength = nVals
	if got := set.Len(); got != wantLength {
		t.Errorf("After adding existing values\n set.Len() = %d want: %d\n",
			got, wantLength)
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
