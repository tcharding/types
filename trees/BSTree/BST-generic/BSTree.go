// Binary Search Tree
//
// Reference: Introduction to Algorithms - Cormen, Leiserson, Rivest, Stein

// Nodes contain a key, satellite data, parent, left child, and right child.

package BSTree

type Node struct {
	// keys are unique across the tree and tree is sorted
	// Formally, for any node x
	// if y is left child of x then y.key < x.key
	// if y is right child of x then y.key > x.key
	key    int
	data   interface{}
	parent *Node // parent Node
	left   *Node // left child
	right  *Node // right child
}

// NewTree: create new empty tree
// NOTE: 'var root Node' will create tree with one node, key value 0
func NewTree() *Node {
	return nil
}

// NewNode: create node with key and data, parent is without parent
func NewNode(key int, data interface{}) *Node {
	return &Node{key, data, nil, nil, nil}
}

// Add x to Node, return true if added
func (n *Node) Add(key int, data interface{}) (*Node, bool) {
	var ok bool
	switch {
	case n == nil:
		return NewNode(key, data), true
	case n.key == key:
		return n, false // unique keys only
	case key < n.key:
		n.left, ok = n.left.Add(key, data)
		// set parent for case when node was created
		if n.left.parent == nil {
			n.left.parent = n
		}
		return n, ok
	case key > n.key:
		n.right, ok = n.right.Add(key, data)
		// set parent for case when node was created
		if n.right.parent == nil {
			n.right.parent = n
		}
		return n, ok
	}
	panic("shouldn't get here")
	return n, false
}

// Size of tree rooted at Node n
func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.Size() + n.right.Size()
}

// Sum of keys in tree rooted at Node n
func (n *Node) Sum() int {
	if n == nil {
		return 0
	}
	return n.key + n.left.Sum() + n.right.Sum()
}

// Height of tree rooted at Node n
func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return 1 + max(n.left.Height(), n.right.Height())
}

// Flatten: in order tree walk of tree rooted at Node n
func (n *Node) Flatten() []int {
	var xs []int

	if n == nil {
		return xs
	}
	if n.left != nil {
		xs = append(xs, n.left.Flatten()...)
	}

	xs = append(xs, n.key)

	if n.right != nil {
		xs = append(xs, n.right.Flatten()...)
	}

	return xs
}

// Max: return maximum key value
func (n *Node) Max() int {
	if n.right == nil {
		return n.key
	}
	return n.right.Max()
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Min: return minimum key value
func (n *Node) Min() int {
	if n.left == nil {
		return n.key
	}
	return n.left.Min()
}

// Elem: true if node with key exists in tree rooted at Node n
func (n *Node) Elem(key int) bool {
	switch {
	case n == nil:
		return false
	case key == n.key:
		return true
	case key < n.key:
		return n.left.Elem(key)
	case key > n.key:
		return n.right.Elem(key)
	}
	panic("shouldn't get here")
	return false
}
