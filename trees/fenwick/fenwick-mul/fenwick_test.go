package fenwick

import "testing"

// tree representing [1, 2, 3, 4, 5]
func testTree() *Fenwick {
	const size = 5
	fen := NewFenwick(size)
	for i := 0; i < size; i++ {
		fen.Update(i, i+1)
	}

	return fen
}

func TestSum(t *testing.T) {
	fen := testTree()

	var tests = []struct {
		index int
		sum   int
	}{
		{0, 1},
		{1, 2},
		{2, 6},
		{3, 24},
		{4, 120},
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
		{1, 2, 4},   // {1, 4, 3, 4, 5}
		{3, 5, 240}, // {1, 4, 3, 20, 5}
	}

	for _, test := range tests {
		fen.Update(test.index, test.val)
		if got := fen.Sum(test.index); got != test.sum {
			t.Errorf("Sum(%d) = %d want: %d\n", test.index, got, test.sum)
		}
	}
}
