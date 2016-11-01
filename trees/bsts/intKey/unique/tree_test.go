package tree

import "testing"

//
//      6
//    /   \
//   4     8
//  / \   / \
// 2   5 7   9
//
func testcase() Tree {
	// sorted := []int{2, 4, 6, 7, 8, 9}
	var tr Tree
	xs := []int{6, 4, 8, 2, 7, 9}
	for _, x := range xs {
		tr.Add(x)
	}
	return tr
}

func TestFlatten(t *testing.T) {
	tr := testcase()
	keys := tr.Flatten()
	want := []int{2, 4, 6, 7, 8, 9}
	if !equal(keys, want) {
		t.Errorf("Fatten failed, keys: %v want: %v\n", keys, want)
	}
}

func TestAdd(t *testing.T) {
	tr := testcase()
	tr.Add(1)
	keys := tr.Flatten()
	want := []int{1, 2, 4, 6, 7, 8, 9}
	if !equal(keys, want) {
		t.Errorf("Fatten failed after initial t.Add(1)\nkeys: %v want: %v\n", keys, want)
	}

	// this tree does not allow multiple same value keys
	tr.Add(1)
	keys = tr.Flatten()
	want = []int{1, 2, 4, 6, 7, 8, 9}
	if !equal(keys, want) {
		t.Errorf("Fatten failed after secondary t.Add(1)\nkeys: %v want: %v\n", keys, want)
	}
}

func TestMinMax(t *testing.T) {
	min := 2
	max := 9
	tr := testcase()

	want := max
	got, err := tr.Max()
	if err != nil {
		t.Errorf("unexpected error: %s\n", err.Error())
	}
	if got != want {
		t.Errorf("tr.Max() = %d want: %d\n", got, want)
	}

	want = min
	got, err = tr.Min()
	if err != nil {
		t.Errorf("unexpected error: %s\n", err.Error())
	}
	if got != want {
		t.Errorf("tr.Min() = %d want: %d\n", got, want)
	}
}

func TestSuccessor(t *testing.T) {
	// []int{2, 4, 6, 7, 8, 9}
	tr := testcase()

	var tests = []struct {
		key  int
		want int
		ok   bool
	}{
		{2, 4, true},
		{6, 7, true},
		{8, 9, true},
		{9, 0, false},
	}

	for _, test := range tests {
		got, ok := tr.Successor(test.key)
		if ok != test.ok {
			t.Errorf("tr.Successor(%d) incorrect ok value\n", test.key)
		}
		if got != test.want {
			t.Errorf("tr.Successor(%d) = %d want: %d\n", test.key, got, test.want)
		}
	}
}

func TestPredecessor(t *testing.T) {
	var tr Tree
	xs := []int{-10, -20, 6, 4, 8, 2, 7, 9}
	for _, x := range xs {
		tr.Add(x)
	}

	var tests = []struct {
		key  int
		want int
		ok   bool
	}{
		{-20, 0, false},
		{-10, -20, true},
		{2, -10, true},
		{4, 2, true},
		{6, 4, true},
		{7, 6, true},
		{8, 7, true},
		{9, 8, true},
	}

	for _, test := range tests {
		got, ok := tr.Predecessor(test.key)
		if ok != test.ok {
			t.Errorf("tr.Predecessor(%d) incorrect ok value\n", test.key)
		}
		if got != test.want {
			t.Errorf("tr.Predecessor(%d) = %d want: %d\n", test.key, got, test.want)
		}
	}
}

func TestDelete(t *testing.T) {
	// []int{2, 4, 6, 7, 8, 9}
	tr := testcase()

	var tests = []struct {
		key int
		ok  bool
		xs  []int // remaining in order slice
	}{
		{6, true, []int{2, 4, 7, 8, 9}},
		{6, false, []int{2, 4, 7, 8, 9}},
		{2, true, []int{4, 7, 8, 9}},
		{9, true, []int{4, 7, 8}},
		{7, true, []int{4, 8}},
		{8, true, []int{4}},
		{4, true, []int{}},
		{8, false, []int{}},
	}

	for _, test := range tests {
		ok := tr.Delete(test.key)
		if ok != test.ok {
			t.Errorf("tr.Delete(%d) incorrect ok value\n", test.key)
		}

		if got := tr.Flatten(); !equal(got, test.xs) {
			t.Errorf("After delete(%d) got: %v want: %v\n", test.key, got, test.xs)
		}
	}

}

func TestSize(t *testing.T) {
	var tr Tree
	xs := []int{3, 0, 5, 1, 2, 4}
	for _, x := range xs {
		tr.Add(x)
	}

	for i := 0; i < 6; i++ {
		want := 6 - i
		size := tr.Size()
		if size != want {
			t.Errorf("tr.Size() = %d want: %d\n", size, want)
		}
		tr.Delete(i)
	}
}

func TestSum(t *testing.T) {
	var tr Tree
	xs := []int{3, 0, 5, 1, 2, 4}
	for _, x := range xs {
		tr.Add(x)
	}
	total := sum(xs)

	for i := 0; i < 6; i++ {
		got := tr.Sum()
		want := total
		if got != want {
			t.Errorf("tr.Size() = %d want: %d\n", got, want)
		}
		tr.Delete(i)
		total -= i
	}
}

func TestElem(t *testing.T) {
	// []int{2, 4, 6, 7, 8, 9}
	tr := testcase()
	var tests = []struct {
		key  int
		want bool
	}{
		{2, true},
		{9, true},
		{6, true},

		{-1, false},
		{5, false},
		{10, false},
	}

	for _, test := range tests {
		if got := tr.Elem(test.key); got != test.want {
			t.Errorf("tr.Elem(%d) failed\n")
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}
