package sset

import "testing"

// test add and len
func Test(t *testing.T) {
	// and, one, that, this
	var tests = []string{"this", "one", "and", "that"}
	set := testSet(tests)

	// test Size
	wantSize := len(tests)
	if got := set.Size(); got != wantSize {
		t.Errorf("set.Size() = %d want: %d\n", got, wantSize)
	}

	// re-add existing values
	for _, s := range tests {
		set.Add(s)
	}
	wantSize = len(tests)
	if got := set.Size(); got != wantSize {
		t.Errorf("After adding existing values\n set.Size() = %d want: %d\nset: %v\n",
			got, wantSize, set)
	}
}

func TestFind(t *testing.T) {
	// and, one, that, this
	var tests = []string{"this", "one", "and", "that"}
	set := testSet(tests)

	// find things that are there
	want := "this"
	got, ok := set.Find(want)
	if got != want {
		t.Errorf("s.Find(%s) = %s want: %s\n", want, got, want)
	}
	if !ok {
		t.Errorf("s.Find(%s) returned spurious not ok", want)
	}

	// find successor
	test := "blue"
	want = "that"
	got, ok = set.Find(want)
	if got != want {
		t.Errorf("s.Find(%s) = %s want: %s\n", test, got, want)
	}
	if !ok {
		t.Errorf("s.Find(%s) returned spurious not ok", test)
	}

	// find false successor
	test = "thiz"
	got, ok = set.Find(test)
	if ok {
		t.Errorf("s.Find(%s) returned spurious ok, got: %s\n", test, got)
	}

}

func TestDelete(t *testing.T) {
	// and, one, that, this
	var tests = []string{"this", "one", "and", "that"}
	set := testSet(tests)
	size := set.Size()

	// delete item and check size
	fn := func(test string) {
		set.Delete(test)
		if got := set.Size(); got != size {
			t.Errorf("set.Delete(%s) error\nset.Size() = %d %v\n", test, got, set)
		}
	}

	// delete non existent
	fn("invalid")

	// delete item from middle of set
	size--
	fn("that")

	// delete item from front of set
	size--
	fn("and")

	// delete item from end of set
	size--
	fn("this")

}

func testSet(v []string) SSet {
	var set SSet
	// test Add
	for _, s := range v {
		set.Add(s)
	}
	return set
}
