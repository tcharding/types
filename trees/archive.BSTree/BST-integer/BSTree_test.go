package BSTree

import "testing"

func TestAdd(t *testing.T) {
	root := Tree()
	root, ok := root.Add(1)
	if !ok {
		t.Errorf("not ok to add initial value\n")
	}

	root, ok = root.Add(1)
	if ok {
		t.Errorf("ok to add repeated value when it should not be\n")
	}
}

func Test(t *testing.T) {
	root := Tree()
	xs := []int{3, 5, 0, 2, 1, 6, 4}
	for _, x := range xs {
		var ok bool
		root, ok = root.Add(x)
		if !ok {
			t.Errorf("not ok to add value: %d\n", x)
		}
	}

	// test Size
	wantSize := len(xs)
	if got := root.Size(); got != wantSize {
		t.Errorf("root.Size() = %d want: %d\n", got, wantSize)
	}

	// test Sum
	wantSum := sum(xs)
	if got := root.Sum(); got != wantSum {
		t.Errorf("root.Sum() = %d want: %d\n", got, wantSum)
	}

	// test Min
	want := 0
	if got := root.Min(); got != want {
		t.Errorf("root.Min() = %d want: %d\n", got, want)
	}

	// test Max
	want = 6
	if got := root.Max(); got != want {
		t.Errorf("root.Max() = %d want: %d\n", got, want)
	}

}

func sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}

func TestHeight(t *testing.T) {
	var tests = []struct {
		xs   []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{0, 1, 2, 3, 4}, 5},
		{[]int{4, 3, 2, 1, 0}, 5},
		{[]int{6, 4, 8, 2, 8, 7, 9}, 3},
	}

	for _, test := range tests {
		root := Tree()
		for _, x := range test.xs {
			root, _ = root.Add(x)
		}
		if got := root.Height(); got != test.want {
			t.Errorf("root.Height() = %d want: %d\n", got, test.want)
			t.Errorf("%v\n", root.Flatten())
		}
	}
}

func TestElem(t *testing.T) {
	root := Tree()
	nVals := 10
	for i := 0; i < nVals; i++ {
		root, _ = root.Add(i)
	}

	// test true
	for i := 0; i < nVals; i++ {
		present := root.Elem(i)
		if !present {
			t.Errorf("root.Elem(%d) false negative\n", i)
		}
	}

	// test false
	for i := nVals + 1; i < nVals*2; i++ {
		present := root.Elem(i)
		if present {
			t.Errorf("root.Elem(%d) false positive\n", i)
		}
	}

}
