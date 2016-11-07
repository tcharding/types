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
	want = []int{1, 1, 2, 4, 6, 7, 8, 9}
	if !equal(keys, want) {
		t.Errorf("Fatten failed after secondary t.Add(1)\nkeys: %v want: %v\n", keys, want)
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

func TestRank(t *testing.T) {
	var tr Tree
	keys := []int{8, 4, 6, 3}
	for _, x := range keys {
		tr.Add(x)
	}

	var tests = []struct {
		input int
		want  int
	}{
		{1, 3},
		{2, 4},
		{3, 6},
		{4, 8},
	}

	for _, test := range tests {

		if got, _ := tr.KthKey(test.input); got != test.want {
			t.Errorf("tr.KthKey(%d) = %d want: %d\n", test.input, got, test.want)
		}
	}
}

func TestMedian(t *testing.T) {
	var tr Tree
	keys := []int{8, 4, 6, 3}
	for _, x := range keys {
		tr.Add(x)
	}

	want := 5
	if got := tr.median(); got != want {
		t.Errorf("meadian: %d want: %d\n", got, want)
	}

	tr.Add(10)
	want = 6
	if got := tr.median(); got != want {
		t.Errorf("meadian: %d want: %d\n", got, want)
	}

	// 3 4 6 8
	var tests = []struct {
		add  int
		want int
	}{
		{5, 5}, // 3 4 5 6 8 10
		{7, 6}, // 3 4 5 6 7 8 10
		{9, 6}, // 3 4 5 6 7 8 9 10
	}
	for _, test := range tests {
		tr.Add(test.add)
		if got := tr.median(); got != test.want {
			t.Errorf("meadian: %d want: %d \ntr: %s\n", got, test.want, tr.String())
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
