// Binary Search Tree, unique integer keys
//
// Reference: Introduction to Algorithms - Cormen, Leiserson, Rivest, Stein
//
// Each node contains a key [and satellite data], each node contains attributes
// left, right, and p that point to nodes corresponding to its left child, right
// child, and parent, respectively. If a child or parent is missing, the
// appropriate attribute contains nil. The root node is the only node in the
// tree whose parent is nil.

// The keys in a binary search tree are always sorted.
// Let x be a node in a binary search tree. If y is a node in the left subtree of
// x, then y.key < x.key. If y is a node in the right subtree of x, then
// y.key > x.key

package tree

import "fmt"

type Tree struct {
	root *Node // May contain multiple same value keys
}

type Node struct {
	// Formally, for any node x
	// if y is left child of x then y.key < x.key
	// if y is right child of x then y.key > x.key
	key    int
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
		return
	}
	// find position
	var p *Node = nil
	n := t.root
	for n != nil {
		p = n
		if new.key == n.key {
			return
		} else if new.key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	// insert node
	if new.key < p.key {
		p.left = new
	} else {
		p.right = new
	}
	new.parent = p
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

// Walk tree in order calling fn for each node
func (n *Node) inorder(fn func(n *Node)) {
	if n != nil {
		n.left.inorder(fn)
		fn(n)
		n.right.inorder(fn)
	}
}

// Max: return maximum key from tree
func (t *Tree) Max() (int, error) {
	if t.root == nil {
		return 0, fmt.Errorf("Max() called on empty tree")
	}
	n := t.root.max()
	return n.key, nil
}

// max: find node with maximum key from tree rooted at n
func (n *Node) max() *Node {
	for n.right != nil {
		n = n.right
	}
	return n
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

// Predecessor: find largest key value smaller than key
// panic if key not present, ok if found
func (t *Tree) Predecessor(key int) (int, bool) {
	n := t.root.find(key)
	if n == nil {
		panic("Predecessor() called with non-existant key")
	}
	prev := n.predecessor()
	if prev == nil {
		return 0, false
	}
	return prev.key, true
}

// predecessor: return node with largest key smaller than n.key
func (n *Node) predecessor() *Node {
	if n.left != nil {
		return n.left.max()
	}
	p := n.parent
	for p != nil && n == p.left {
		n = p
		p = p.parent
	}
	return p
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

// Sum keys in tree
func (t *Tree) Sum() int {
	if t.root == nil {
		return 0
	}
	return t.root.sum()
}

// Sum of keys in tree rooted at node n
func (n *Node) sum() int {
	total := 0
	fn := func(n *Node) {
		if n != nil {
			total += n.key
		}
	}
	n.inorder(fn)
	return total

}

// Height: longest path from root of tree to any leaf
func (t *Tree) Height() int {
	return t.root.height()
}

// height of tree rooted at node n
func (n *Node) height() int {
	if n == nil {
		return 0
	}
	return 1 + max(n.left.height(), n.right.height())
}

// Elem: true if node with key exists in tree
func (t *Tree) Elem(key int) bool {
	return t.root.elem(key)
}

// elem: true if node with key exists in tree rooted at node n, iterative version
func (n *Node) elem(key int) bool {
	found := false
	for {
		if n == nil {
			break // found = false
		}
		if n.key == key {
			found = true
			break
		}
		if key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return found
}

// // elem: true if node with key exists in tree rooted at node n, recursive version
// func (n *Node) elem(key int) bool {
// 	if n == nil {
// 		return false
// 	}
// 	if key == n.key {
// 		return true
// 	} else if key < n.key {
// 		return n.left.elem(key)
// 	}
// 	return n.right.elem(key)
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
