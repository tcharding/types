package bag

import "testing"

// len, add, delete, find

func Test(t *testing.T) {
	bag := make(Bag)
	nVals := 5

	// test Add
	for i := 0; i < nVals; i++ {
		bag.Add(i)
	}

	// test Len
	wantLength := nVals
	if got := bag.Len(); got != wantLength {
		t.Errorf("bag.Len() = %d want: %d\n", got, wantLength)
	}

	// re-add existing values
	for i := 0; i < nVals; i++ {
		bag.Add(i)
	}
	wantLength = nVals * 2
	if got := bag.Len(); got != wantLength {
		t.Errorf("After adding existing values\n bag.Len() = %d want: %d\n",
			got, wantLength)
	}

	// test Find
	for i := 0; i < nVals; i++ {
		x, ok := bag.Find(i)
		if !ok || x != 2 {
			t.Errorf("bag.Find() failed to find value: %d got: %d\n", i, x)
		}
	}

	for i := nVals + 1; i < nVals*2; i++ {
		_, found := bag.Find(i)
		if found != false {
			t.Errorf("bag.Find() failed to find value: %d\n", i)
		}
	}

}
