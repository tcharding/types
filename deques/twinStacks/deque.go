// dequeue implemented using two stacks
package deque

import "fmt"

type Deque struct {
	front stack
	back  stack
}

type stack []int

// push a new integer onto the stack
func (s *stack) push(x int) {
	*s = append(*s, x)
}

// pop: remove and return top element of stack, return false if stack is empty
func (s *stack) pop() (int, bool) {
	if s.len() == 0 {
		return 0, false
	}

	i := len(*s) - 1
	x := (*s)[i]
	*s = (*s)[:i]

	return x, true
}

func (s *stack) len() int {
	return len(*s)
}

// enqueue front
func (d *Deque) EnqueueF(x int) {
	d.front.push(x)
	d.balance()
}

// enqueue back
func (d *Deque) EnqueueB(x int) {
	d.back.push(x)
	d.balance()
}

// dequeue front
func (d *Deque) DequeueF() (int, bool) {
	var x int
	var ok bool
	if d.front.len() == 0 {
		x, ok = d.back.pop()
	} else {
		x, ok = d.front.pop()
	}
	if !ok {
		return 0, false
	}
	d.balance()
	return x, true
}

// dequeue back
func (d *Deque) DequeueB() (int, bool) {
	var x int
	var ok bool
	if d.back.len() == 0 {
		x, ok = d.front.pop()
	} else {
		x, ok = d.back.pop()
	}
	if !ok {
		return 0, false
	}
	d.balance()
	return x, true
}

// balance stacks if needed
func (d *Deque) balance() {
	small, big := order(&d.front, &d.back)
	if small.len() == 0 && big.len() == 1 {
		return
	}
	if 3*d.front.len() < d.back.len() ||
		3*d.back.len() < d.front.len() {
		d.rebalance()
	}
}

// rebalance stacks
func (d *Deque) rebalance() {
	small, big := order(&d.front, &d.back)
	half := (small.len() + big.len()) / 2
	tmpB := &stack{}
	tmpS := &stack{}
	mvN(tmpB, big, half) // store half of big
	mvAll(tmpS, small)   // store small
	mvAll(small, big)    // put bottom half of big onto small
	mvAll(small, tmpS)   // restore small
	mvAll(big, tmpB)     // restore big
}

func order(a, b *stack) (smallest, bigest *stack) {
	if a.len() < b.len() {
		return a, b
	}
	return b, a
}

// mvN: move n items from dst and push to src
func mvN(dst, src *stack, n int) {
	for i := 0; i < n; i++ {
		x, _ := src.pop()
		dst.push(x)
	}
}

// mvAll: move all items from src to dst

func mvAll(dst, src *stack) {
	mvN(dst, src, src.len())
}

func (d *Deque) String() string {
	front := fmt.Sprintf("[%v]", d.front)
	back := fmt.Sprintf("[%v]", d.back)
	return fmt.Sprintf("%s | %s", back, front)
}
