// Implements list interface
//
// type List interface {
// 	Len() int
// 	Get(i int) interface{}
// 	Set(i int, x interface{})
// 	Insert(i int, x interface{})
// 	Delete(i int) interface{}
// }

// Initial attempt at implementing list interface. Inefficient, useful for
// testing against more complex implementations.

package list

type List []int

// Len: length of list
func (l *List) Len() int {
	return len(*l)
}

// Get item at index i
func (l *List) Get(i int) int {
	return (*l)[i]
}

// Set item at index i to x
func (l *List) Set(i, x int) {
	(*l)[i] = x
}

// Insert x at index i
func (l *List) Insert(i, x int) {
	if l.Len() != i {
		panic("insert at end of list only")
	}
	*l = append(*l, x)
}

// Delete item at index i
func (l *List) Delete(i int) {
	*l = append((*l)[:i], (*l)[i+1:]...)
}
