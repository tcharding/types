// minimum heap of integers
package heap

import "fmt"

type heap struct {
	xs []int
}

// MakeHeap: make a minimum heap from values in xs, O(n)
func MakeHeap(xs []int) *heap {
	var h heap
	h.xs = make([]int, len(xs)+1)
	for i, x := range xs {
		h.xs[i+1] = x
	}
	for i := len(xs) / 2; i >= 1; i-- {
		h.bubbleDown(i)

	}
	return &h
}

// // MakeHeap: make a minimum heap from values in xs, O(nlog(n))
// func MakeHeap(xs []int) *heap {
// 	h := heap{[]int{0}} // heap is 1-indexed
// 	for _, x := range xs {
// 		h.Insert(x)
// 	}
// 	return &h
// }

func (h *heap) Len() int {
	return len(h.xs) - 1 // heap is 1-indexed
}

// Insert x into the heap
func (h *heap) Insert(x int) {
	(*h).xs = append(h.xs, x)
	h.bubbleUp(len(h.xs) - 1)
}

func (h *heap) bubbleUp(k int) {
	p, ok := parent(k)
	if !ok {
		return // k is root node
	}
	if h.xs[p] > h.xs[k] {
		h.xs[k], h.xs[p] = h.xs[p], h.xs[k]
		h.bubbleUp(p)
	}
}

// Min: get the minimum value of the heap
// do not remove value from heap
func (h *heap) Min() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return h.xs[1], true
}

// ExtractMin: get minimum value of heap
// and remove value from heap
func (h *heap) ExtractMin() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	v := h.xs[1]
	h.xs[1] = h.xs[h.Len()]
	(*h).xs = h.xs[:h.Len()]
	h.bubbleDown(1)
	return v, true
}

func (h *heap) bubbleDown(k int) {
	min := k
	c := left(k)

	// find index of minimum value (k, k's left child, k's right child)
	for i := 0; i < 2; i++ {
		if (c + i) <= h.Len() {
			if h.xs[min] > h.xs[c+i] {
				min = c + i
			}
		}
	}
	if min != k {
		h.xs[k], h.xs[min] = h.xs[min], h.xs[k]
		h.bubbleDown(min)
	}
}

func (h *heap) String() string {
	return fmt.Sprintf("MIN-HEAP %v", h.xs)
}

// get index of parent of node at index k
func parent(k int) (int, bool) {
	if k == 1 {
		return 0, false
	}
	return k / 2, true
}

// get index of left child of node at index k
func left(k int) int {
	return 2 * k
}

// get index of right child of node at index k
func right(k int) int {
	return 2*k + 1
}
