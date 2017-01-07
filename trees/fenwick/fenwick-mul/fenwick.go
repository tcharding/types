// Fenwick tree, update is done via multiplication.
package fenwick

// ref: https://en.wikipedia.org/wiki/Fenwick_tree

type Fenwick struct {
	tree []int
}

// NewFenwick: Build Fenwick tree to hold n values.
func NewFenwick(n int) *Fenwick {
	fen := &Fenwick{
		tree: make([]int, n),
	}
	for i := range fen.tree {
		fen.tree[i] = 1
	}
	return fen
}

// Sum all values upto and including index.
func (fen *Fenwick) Sum(index int) int {
	sum := 1
	index++
	for index > 0 {
		sum *= fen.tree[index-1]
		index -= lsb(index)
	}
	return sum
}

// lsb: Least Significant Bit
func lsb(x int) int {
	return x & -x
}

// Update using addition index by value.
func (fen *Fenwick) Update(index, value int) {
	index++
	for index <= fen.Size() {
		fen.tree[index-1] *= value
		index += lsb(index)
	}
}

// Size: Number of values stored by tree.
func (fen *Fenwick) Size() int {
	return len(fen.tree)
}
