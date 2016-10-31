// Node can be used as a stand alone tree or to construct more complex trees
// Each node contains integer data (key).
//
// Formally, for any node n:
//  n.key > n.left.key
//  n.key < n.right.key

package Node

//
// WARNING: the zero value of Node is a node with key = 0 and all attributes set to nil
//

type Node struct {
	key    int
	parent *Node
	left   *Node
	right  *Node
}

// Tree: return new tree
// this is useful so as not to create a node with initial key value of 0
func Tree() *Node {
	return nil
}

// Add key to Node, return true if added
func (n *Node) Add(key int) (*Node, bool) {
	if n == nil {
		return &Node{key, nil, nil, nil}, true
	}
	var ok bool
	if key < n.key {
		n.left, ok = n.left.Add(key)
		if n.left.parent == nil {
			n.left.parent = n
		}
		return n, ok
	} else {
		n.right, ok = n.right.Add(key)
		if n.right.parent == nil {
			n.right.parent = n
		}
		return n, ok
	}
	panic("shouldn't get here")
	return n, false
}

// size of tree rooted at n
func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.Size() + n.right.Size()
}

// sum of keys of tree rooted at n
func (n *Node) Sum() int {
	if n == nil {
		return 0
	}
	return n.key + n.left.Sum() + n.right.Sum()
}

// height of tree rooted at n
func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return 1 + max(n.left.Height(), n.right.Height())
}

// flatten in-order keys of tree rooted at n
func (n *Node) Flatten() []int {
	var keys []int
	if n == nil {
		return keys
	}
	if n.left != nil {
		keys = append(keys, n.left.Flatten()...)
	}
	keys = append(keys, n.key)
	if n.right != nil {
		keys = append(keys, n.right.Flatten()...)
	}
	return keys
}

// max key of tree rooted at n
func (n *Node) Max() int {
	if n.right == nil {
		return n.key
	}
	return n.right.Max()
}

func max(a, b int) int {
	switch {
	case a == b:
		return a
	case a < b:
		return b
	case a > b:
		return a
	default:
		panic("shouldn't get here")
	}
}

// min key of tree rooted at n
func (n *Node) Min() int {
	if n.left == nil {
		return n.key
	}
	return n.left.Min()
}

// true if key is in tree rooted at n
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
