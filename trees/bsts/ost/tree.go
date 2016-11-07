// Ordered Statistical Tree - binary search tree with rank
package tree

import (
	"bytes"
	"fmt"
	"math"
)

type Tree struct {
	root *Node // May contain multiple same value keys
}

type Node struct {
	// Formally, for any node x
	// if y is left child of x then y.key < x.key
	// if y is right child of x then y.key > x.key
	key    int
	rank   int
	parent *Node // parent node
	left   *Node // left child
	right  *Node // right child
}

// Add key to tree
func (t *Tree) Add(key int) {
	n := newNode(key)
	t.insert(n)
}

// newNode: create a new node from key
func newNode(key int) *Node {
	var n Node
	n.key = key
	return &n
}

// insert node n into t
func (t *Tree) insert(new *Node) {
	if t.root == nil {
		t.root = new
		new.rank = 1 // 1 indexed
		return
	}
	// find position
	var p *Node = nil
	n := t.root
	for n != nil {
		p = n
		if new.key < n.key {
			n = n.left
			p.rank++
			p.right.incRank()
		} else {
			n = n.right
		}
	}
	// insert node
	if new.key < p.key {
		p.left = new
		new.rank = p.rank - 1
	} else {
		p.right = new
		new.rank = p.rank + 1
	}
	new.parent = p
}

// incRank: increment rank in all nodes rooted at n
func (n *Node) incRank() {
	fn := func(n *Node) {
		n.rank++
	}
	n.inorder(fn)
}

// Successor: find smallest key value larger than key
// panic if key not present, ok if found
func (t *Tree) Successor(key int) (int, bool) {
	n := t.root.find(key)
	if n == nil {
		panic("Succesor() called with non-existant key")
	}
	next := n.successor()
	if next == nil {
		return 0, false
	}
	return next.key, true
}

// find node by key
func (n *Node) find(key int) *Node {
	for n != nil && key != n.key {
		if key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return n
}

// successor: return node with smallest key larger than n.key
func (n *Node) successor() *Node {
	if n.right != nil {
		return n.right.min()
	}
	p := n.parent
	for p != nil && n == p.right {
		n = p
		p = p.parent
	}
	return p
}

// Min: return minimum key from tree
func (t *Tree) Min() (int, error) {
	if t.root == nil {
		return 0, fmt.Errorf("Min() called on empty tree")
	}
	n := t.root.min()
	return n.key, nil
}

// min: find node with minimum key from tree rooted at n
func (n *Node) min() *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

// Size of tree
func (t *Tree) Size() int {
	if t.root == nil {
		return 0
	}
	return t.root.size()
}

// size of tree rooted at node n
func (n *Node) size() int {
	total := 0
	fn := func(n *Node) {
		if n != nil {
			total++
		}
	}
	n.inorder(fn)
	return total
}

// Walk tree in order calling fn for each node
func (n *Node) inorder(fn func(n *Node)) {
	if n != nil {
		n.left.inorder(fn)
		fn(n)
		n.right.inorder(fn)
	}
}

// kthKey: return kth key index starts at 1
func (t *Tree) KthKey(rank int) (int, bool) {
	if t.root == nil {
		return 0, false
	}
	n, ok := t.root.kthKey(rank)
	if !ok {
		return 0, false
	}
	return n.key, true
}

func (n *Node) kthKey(rank int) (*Node, bool) {
	if n == nil {
		return nil, false
	}
	if rank == n.rank {
		return n, true
	}
	if rank < n.rank {
		return n.left.kthKey(rank)
	}
	return n.right.kthKey(rank)
}

func (t *Tree) String() string {
	var buf bytes.Buffer
	fn := func(n *Node) {
		s := fmt.Sprintf("(%d %d) ", n.key, n.rank)
		buf.WriteString(s)
	}
	t.root.inorder(fn)
	return buf.String()
}

// calculate the median key value
// if t.size is even return average of 2 median values
func (t *Tree) median() int {
	size := t.Size()
	if isOdd(size) {
		k := size/2 + 1
		key, _ := t.KthKey(k)
		return key
	}
	k := size / 2
	n, _ := t.root.kthKey(k)
	sum := n.key + n.successor().key
	return sum / 2
}

func isOdd(x int) bool {
	return math.Mod(float64(x), 2) == 1
}

// transplant replaces subtree rooted at u with subtree rooted at v
func (t *Tree) transplant(u, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

// Flatten: return inorder slice of keys
func (t *Tree) Flatten() []int {
	var keys []int
	fn := func(n *Node) {
		keys = append(keys, n.key)
	}
	t.root.inorder(fn)
	return keys
}

// Delete node with key from tree
// true if deleted
func (t *Tree) Delete(key int) bool {
	n := t.root.find(key)
	if n == nil {
		return false
	}
	t.delete(n)
	return true
}

func (t *Tree) delete(d *Node) {
	if d.left == nil {
		t.transplant(d, d.right)
	} else if d.right == nil {
		t.transplant(d, d.left)
	} else {
		n := d.right.min()
		if n.parent != d {
			t.transplant(n, n.right)
			n.right = d.right
			n.right.parent = n
		}
		t.transplant(d, n)
		n.left = d.left
		n.left.parent = n
	}
}
