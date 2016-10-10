// Doubly linked list
//
// Dynamic data set supporting the following operations
//
// delete: O(1)
// search: O(n)
// insert: O(1) (head, tail, or after node)

package linkedList

// type ListNode interface {
// 	Equal(n ListNode) bool
// 	Next() ListNode
// 	Prev() ListNode
// }

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	value interface{}
	next  *Node
	prev  *Node
}

func (p *LinkedList) Length() int {
	return p.length
}

func (p *LinkedList) AddHead(value interface{}) {
	n := Node{value, p.head, p.tail}
	if p.head == nil {
		n.prev = &n
		n.next = &n
		p.head = &n
		p.tail = &n

	} else {
		p.head.prev = &n
		p.tail.next = &n
		p.head = &n
	}
	p.length++
}

func (p *LinkedList) AddTail(value interface{}) {
	n := Node{value, p.head, p.tail}
	if p.head == nil {
		n.prev = &n
		n.next = &n
		p.head = &n
		p.tail = &n

	} else {
		p.head.prev = &n
		p.tail.next = &n
		p.tail = &n
	}
	p.length++
}

// ForEach: call func once for each element of list
func (p *LinkedList) ForEach(f func(n *Node)) {
	np := p.head
	if np == nil {
		return
	}
	for {
		f(np)

		if np == p.tail {
			break
		}
		np = np.next
	}
}

// Delete n from list
func (p *LinkedList) Delete(n *Node) {
	if n == nil {
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	if p.head == n {
		p.head = n.next
	}
	if p.tail == n {
		p.tail = n.prev
	}
	p.length--
}

func (p *LinkedList) Get(f func(n *Node) bool) *Node {
	np := p.head
	if np == nil {
		return nil
	}
	for {
		if f(np) {
			return np
		}
		if np == p.tail {
			break
		}
	}
	return nil
}
