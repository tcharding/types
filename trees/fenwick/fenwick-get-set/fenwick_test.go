package fenwick

import "testing"

// tree representing [1, 2, 3, 4, 5]
func testTree() *Fenwick {
	data := []int{1, 2, 3, 4, 5}
	fen := NewFenwick(len(data))
	for i, x := range data {
		fen.Set(i, x)
	}
	return fen
}

func TestSet(t *testing.T) {
	fen := testTree()

	var tests = []struct {
		index         int // NOTE: don't test last index.
		value         int
		wantThisIndex int
		wantNextIndex int
	}{
		{1, 6, 7, 10},   // {1, 6, 3, 4, 5}
		{3, 10, 20, 25}, // {1, 6, 3, 10, 5}
	}

	for _, test := range tests {
		fen.Set(test.index, test.value)
		if got := fen.Get(test.index); got != test.value {
			t.Errorf("Set failed")
		}

		if got := fen.Sum(test.index); got != test.wantThisIndex {
			t.Errorf("Sum(%d) = %d want: %d\n", test.index, got, test.wantThisIndex)
		}

		if got := fen.Sum(test.index + 1); got != test.wantNextIndex {
			t.Errorf("Sum(%d) = %d want: %d\n", test.index+1, got, test.wantThisIndex)
		}

	}
}

func TestGet(t *testing.T) {
	fen := testTree()

	var tests = []struct {
		index int
		want  int
	}{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}

	for _, test := range tests {
		if got := fen.Get(test.index); got != test.want {
			t.Errorf("Get(%d) = %d want: %d\n", test.index, got, test.want)
		}
	}
}

func TestSum(t *testing.T) {
	fen := testTree()

	var tests = []struct {
		index int
		sum   int
	}{
		{0, 1},
		{1, 3},
		{2, 6},
		{3, 10},
		{4, 15},
	}

	for _, test := range tests {
		if got := fen.Sum(test.index); got != test.sum {
			t.Errorf("Sum(%d) = %d want: %d\n", test.index, got, test.sum)
		}
	}
}

func TestUpdate(t *testing.T) {
	fen := testTree()
	var tests = []struct {
		index int
		val   int
		sum   int
	}{
		{1, 1, 4},  // {1, 3, 3, 4, 5}
		{3, 1, 12}, // {1, 3, 3, 5, 5}
	}

	for _, test := range tests {
		fen.Update(test.index, test.val)
		if got := fen.Sum(test.index); got != test.sum {
			t.Errorf("Sum(%d) = %d want: %d\n", test.index, got, test.sum)
		}
	}
}
