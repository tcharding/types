package BSTree

import (
	"fmt"
	"os"
	"testing"
)

func TestAddBasic(t *testing.T) {
	var tr Tree
	tr.Add(1, nil)
}

// test test file utilities
func TestUtils(t *testing.T) {
	var tr Tree
	tr.Add(1, nil)
	// tr := testTree()
	fmt.Fprintf(os.Stderr, "tr: %v\n", tr)
	keys := inorderKeySlice(tr)
	// want := []int{2, 4, 6, 7, 8, 9}
	want := []int{1}
	if !equal(keys, want) {
		t.Errorf("test utilities failed, keys: %v want: %v\n", keys, want)
	}
}

//
//      6
//    /   \
//   4     8
//  / \   / \
// 2   5 7   9
//
func testTree() Tree {
	var tr Tree
	xs := []int{6, 4, 8, 2, 7, 9}
	for _, x := range xs {
		tr.Add(x, nil)
	}
	return tr
}

func inorderKeySlice(tr Tree) []int {
	var keys []int
	fn := func(n *Node) {
		fmt.Fprintf(os.Stderr, "adding: %d\n", n.key)
		keys = append(keys, n.key)
	}
	inorderTreeWalk(tr.root, fn)
	return keys
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

// func TestPrintTree(t *testing.T) {
// 	tr := testTree()
// 	out := new(bytes.Buffer)
// 	tr.Write(out)
// 	want := "[ 2 4 5 6 7 8 9 ]"
// 	if got := out.String(); got != want {
// 		t.Errorf("tr.Print() = %s\n         want: %s\n", got, want)
// 	}
// }

// func Test(t *testing.T) {
// 	root := NewTree()
// 	xs := []int{3, 5, 0, 2, 1, 6, 4}
// 	for _, x := range xs {
// 		var ok bool
// 		root, ok = root.Add(x, nil)
// 		if !ok {
// 			t.Errorf("not ok to add value: %d\n", x)
// 		}
// 	}

// 	// test Size
// 	wantSize := len(xs)
// 	if got := root.Size(); got != wantSize {
// 		t.Errorf("root.Size() = %d want: %d\n", got, wantSize)
// 	}

// 	// test Sum
// 	wantSum := sum(xs)
// 	if got := root.Sum(); got != wantSum {
// 		t.Errorf("root.Sum() = %d want: %d\n", got, wantSum)
// 	}

// 	// test Min
// 	want := 0
// 	if got := root.Min(); got != want {
// 		t.Errorf("root.Min() = %d want: %d\n", got, want)
// 	}

// 	// test Max
// 	want = 6
// 	if got := root.Max(); got != want {
// 		t.Errorf("root.Max() = %d want: %d\n", got, want)
// 	}

// }

// func sum(xs []int) int {
// 	total := 0
// 	for _, x := range xs {
// 		total += x
// 	}
// 	return total
// }

// func TestHeight(t *testing.T) {
// 	var tests = []struct {
// 		xs   []int
// 		want int
// 	}{
// 		{[]int{1}, 1},
// 		{[]int{0, 1, 2, 3, 4}, 5},
// 		{[]int{4, 3, 2, 1, 0}, 5},
// 		{[]int{6, 4, 8, 2, 8, 7, 9}, 3},
// 	}

// 	for _, test := range tests {
// 		root := NewTree()
// 		for _, x := range test.xs {
// 			root, _ = root.Add(x, nil)
// 		}
// 		if got := root.Height(); got != test.want {
// 			t.Errorf("root.Height() = %d want: %d\n", got, test.want)
// 			t.Errorf("%v\n", root.Flatten())
// 		}
// 	}
// }

// func TestElem(t *testing.T) {
// 	root := NewTree()
// 	nVals := 10
// 	for i := 0; i < nVals; i++ {
// 		root, _ = root.Add(i, nil)
// 	}

// 	// test true
// 	for i := 0; i < nVals; i++ {
// 		present := root.Elem(i)
// 		if !present {
// 			t.Errorf("root.Elem(%d) false negative\n", i)
// 		}
// 	}

// 	// test false
// 	for i := nVals + 1; i < nVals*2; i++ {
// 		present := root.Elem(i)
// 		if present {
// 			t.Errorf("root.Elem(%d) false positive\n", i)
// 		}
// 	}

// }

// func TestSuccessor(t *testing.T) {
// 	root := NewTree()
// 	xs := []int{6, 4, 8, 2, 5, 7, 9}
// 	for _, x := range xs {
// 		root, _ = root.Add(x, nil)
// 	}

// 	var tests = []struct {
// 		key  int
// 		want int
// 	}{
// 		{8, 9},
// 		{7, 8},
// 		{5, 6},
// 	}

// 	for _, test := range tests {
// 		got, ok := root.Successor(test.key)
// 		if !ok {
// 			t.Errorf("spurious not ok")
// 		}
// 		if got != test.want {
// 			t.Errorf("root.Successor(%d) = %d want: %d\n", test.key, got, test.want)
// 		}
// 	}
// }

// func TestPredecessor(t *testing.T) {
// 	root := NewTree()
// 	xs := []int{6, 4, 8, 2, 5, 7, 9}
// 	for _, x := range xs {
// 		root, _ = root.Add(x, nil)
// 	}

// 	var tests = []struct {
// 		key  int
// 		want int
// 	}{
// 		{5, 4},
// 		{7, 6},
// 		{9, 8},
// 	}

// 	for _, test := range tests {
// 		got, ok := root.Predecessor(test.key)
// 		if !ok {
// 			t.Errorf("spurious not ok")
// 		}
// 		if got != test.want {
// 			t.Errorf("root.Predecessor(%d) = %d want: %d\n", test.key, got, test.want)
// 		}
// 	}
// 	_, ok := root.Predecessor(2)
// 	if ok {
// 		t.Errorf("spurious ok")
// 	}
// }
