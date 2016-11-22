// N-ary tree of integers, with no explicit order.
package tree

import (
	"bytes"
	"fmt"
)

type Node struct {
	key         int
	firstChild  *Node
	nextSibling *Node
}

// Create a new node with key k
func NewNode(k int) *Node {
	return &Node{k, nil, nil}
}

// AddChild: Add child node with key k to n,
// return newly created child node.
func (n *Node) AddChild(k int) *Node {
	child := Node{k, nil, nil}
	n.AddChildNode(&child)
	return &child
}

// AddChildNode: Add child to n.
func (n *Node) AddChildNode(child *Node) {
	(*child).nextSibling = n.firstChild
	(*n).firstChild = child
}

// Size of tree rooted at n.
func (n *Node) Size() int {
	size := 1
	for p := n.firstChild; p != nil; p = p.nextSibling {
		size += p.Size()
	}
	return size
}

// String: Convert n to string format.
func (n *Node) String() string {
	var (
		buf      bytes.Buffer
		depth    int
		toString = func(n *Node, depth int) {}
	)

	toString = func(n *Node, depth int) {
		if n == nil {
			return
		}
		bufferSpaces(&buf, depth)
		buf.WriteString(fmt.Sprintf("%d", n.key))
		buf.WriteString("\n")
		for cp := n.firstChild; cp != nil; cp = cp.nextSibling {
			toString(cp, depth+1)
		}
	}

	depth = 0
	toString(n, depth)
	return buf.String()
}

func bufferSpaces(buf *bytes.Buffer, n int) {
	for i := 0; i < n; i++ {
		buf.WriteString("  ")
	}
}
