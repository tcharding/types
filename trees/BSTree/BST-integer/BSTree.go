// Binary Search Tree - Unique Integer tree

// Data is held in each node, leaf nodes have two nil children.

package BSTree

type Node struct {
	x      int
	parent *Node
	left   *Node
	right  *Node
}

// Tree: return new tree
// this is needed so as not to create a node with initial x value of 0
func Tree() *Node {
	return nil
}

// Add x to Node, return true if added
func (n *Node) Add(x int) (*Node, bool) {
	if n == nil {
		return &Node{x, nil, nil, nil}, true
	}

	if n.x == x {
		return n, false
	}

	var ok bool
	if x < n.x {
		n.left, ok = n.left.Add(x)
		if n.left.parent == nil {
			n.left.parent = n
		}
		return n, ok
	} else if x > n.x {
		n.right, ok = n.right.Add(x)
		if n.right.parent == nil {
			n.right.parent = n
		}
		return n, ok
	}

	panic("shouldn't get here")
	return n, false
}

func (n *Node) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.Size() + n.right.Size()
}

func (n *Node) Sum() int {
	if n == nil {
		return 0
	}
	return n.x + n.left.Sum() + n.right.Sum()
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return 1 + max(n.left.Height(), n.right.Height())
}

func (n *Node) Flatten() []int {
	var xs []int

	if n == nil {
		return xs
	}
	if n.left != nil {
		xs = append(xs, n.left.Flatten()...)
	}

	xs = append(xs, n.x)

	if n.right != nil {
		xs = append(xs, n.right.Flatten()...)
	}

	return xs
}

func (n *Node) Max() int {
	if n.right == nil {
		return n.x
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

func (n *Node) Min() int {
	if n.left == nil {
		return n.x
	}
	return n.left.Min()
}

func (n *Node) Elem(x int) bool {
	switch {
	case n == nil:
		return false
	case x == n.x:
		return true
	case x < n.x:
		return n.left.Elem(x)
	case x > n.x:
		return n.right.Elem(x)
	}

	panic("shouldn't get here")
	return false
}
