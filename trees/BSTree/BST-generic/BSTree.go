// Binary Search Tree
//
// Reference: Introduction to Algorithms - Cormen, Leiserson, Rivest, Stein
//
// Each node contains a key and satellite data, each node contains attributes
// left, right, and p that point to nodes corresponding to its left child, right
// child, and parent, respectively. If a child or parent is missing, the
// appropriate attribute contains nil. The root node is the only node in the
// tree whose parent is nil.

// The keys in a binary search tree are always sorted.
// Let x be a node in a binary search tree. If y is a node in the left subtree of
// x, then y.key <= x.key. If y is a node in the right subtree of x, then
// y.key >= x.key

package BSTree

type Node struct {
	// Multiple keys of same value permitted
	// Formally, for any node x
	// if y is left child of x then y.key < x.key
	// if y is right child of x then y.key => x.key
	key    int
	data   interface{}
	parent *Node // parent node
	left   *Node // left child
	right  *Node // right child
}

type Tree struct {
	root *Node
}

// Walk tree in order, call fn for each node
func (t Tree) inorder(n *Node, fn func(n *Node)) {
	if n != nil {
		inorderTreeWalk(n.left, fn)
		fn(n)
		inorderTreeWalk(n.right, fn)
	}
}

// func (t Tree) Add(key int, data interface{}) {
// 	n := newNode(key, data)
// 	t.treeInsert(n)
// }

// func (t Tree) Write(w io.Writer) {
// 	var buf bytes.Buffer
// 	fn := func(n *Node) {
// 		buf.WriteString(fmt.Sprintf("%d ", n.key))
// 	}

// 	buf.WriteString("[ ")
// 	inorderTreeWalk(t.root, fn)
// 	buf.WriteRune(']')
// 	fmt.Fprintf(w, "%v\n", buf)
// }

// // newNode: create node with key and data, node has no parent
// func newNode(key int, data interface{}) *Node {
// 	return &Node{key, data, nil, nil, nil}
// }

// // Add x to node, return true if added
// func (n *Node) Add(key int, data interface{}) (*Node, bool) {
// 	var ok bool
// 	switch {
// 	case n == nil:
// 		return NewNode(key, data), true
// 	case n.key == key:
// 		return n, false // unique keys only
// 	case key < n.key:
// 		n.left, ok = n.left.Add(key, data)
// 		// set parent for case when node was created
// 		if n.left.parent == nil {
// 			n.left.parent = n
// 		}
// 		return n, ok
// 	case key > n.key:
// 		n.right, ok = n.right.Add(key, data)
// 		// set parent for case when node was created
// 		if n.right.parent == nil {
// 			n.right.parent = n
// 		}
// 		return n, ok
// 	}
// 	panic("shouldn't get here")
// 	return n, false
// }

// // Size of tree rooted at node n
// func (n *Node) Size() int {
// 	if n == nil {
// 		return 0
// 	}
// 	return 1 + n.left.Size() + n.right.Size()
// }

// // Sum of keys in tree rooted at node n
// func (n *Node) Sum() int {
// 	if n == nil {
// 		return 0
// 	}
// 	return n.key + n.left.Sum() + n.right.Sum()
// }

// // Height of tree rooted at node n
// func (n *Node) Height() int {
// 	if n == nil {
// 		return 0
// 	}
// 	return 1 + max(n.left.Height(), n.right.Height())
// }

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// // Flatten: in order tree walk of tree rooted at node n
// func (n *Node) Flatten() []int {
// 	var xs []int

// 	if n == nil {
// 		return xs
// 	}
// 	if n.left != nil {
// 		xs = append(xs, n.left.Flatten()...)
// 	}

// 	xs = append(xs, n.key)

// 	if n.right != nil {
// 		xs = append(xs, n.right.Flatten()...)
// 	}

// 	return xs
// }

// // Max: return maximum key value
// func (n *Node) Max() int {
// 	for n.right != nil {
// 		n = n.right
// 	}
// 	return n.key
// }

// // Min: return minimum key value
// func (n *Node) Min() int {
// 	for n.left != nil {
// 		n = n.left
// 	}
// 	return n.key
// }

// // Elem: true if node with key exists in tree rooted at node n
// func (n *Node) Elem(key int) bool {
// 	for {
// 		if n == nil {
// 			return false
// 		}
// 		if n.key == key {
// 			return true
// 		}
// 		if key < n.key {
// 			n = n.left
// 		} else {
// 			n = n.right
// 		}
// 	}
// 	panic("shouldn't get here")
// 	return false
// }

// // Successor: return smallest key from tree rooted at node n bigger than key
// func (n *Node) Successor(key int) (int, bool) {
// 	x, ok := n.get(key)
// 	if !ok {
// 		panic("cannot call successor with a non-existent key")
// 	}
// 	if x.right != nil {
// 		return x.right.Min(), true
// 	}
// 	y := x.parent
// 	for {
// 		if y == nil {
// 			return 0, false
// 		}
// 		if x == y.left {
// 			return y.key, true
// 		}
// 		x = y
// 		y = y.parent
// 	}
// 	panic("shouldn't get here")
// 	return 0, false
// }

// // Predecessor: return biggest key from tree rooted at node n smaller than key
// func (n *Node) Predecessor(key int) (int, bool) {
// 	x, ok := n.get(key)
// 	if !ok {
// 		panic("cannot call predecessor with a non-existent key")
// 	}
// 	if x.left != nil {
// 		return x.left.Max(), true
// 	}

// 	y := x.parent
// 	for {
// 		if y == nil {
// 			return 0, false
// 		}
// 		if x == y.right {
// 			return y.key, true
// 		}
// 		x = y
// 		y = x.parent
// 	}
// 	return 0, false
// }

// // get node for key
// func (n *Node) get(key int) (*Node, bool) {
// 	for {
// 		if n == nil {
// 			return nil, false
// 		}
// 		if n.key == key {
// 			return n, true
// 		}
// 		if key < n.key {
// 			n = n.left
// 		} else {
// 			n = n.right
// 		}
// 	}
// 	panic("shouldn't get here")
// 	return nil, false
// }
