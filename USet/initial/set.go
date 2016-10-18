// USet: Un-ordered set of Integers, implements the USet interface

// Initial attempt at implementing list interface. Inefficient, useful for
// testing against more complex implementations.

package uset

type USet map[int]bool

func (s *USet) Len() int {
	return len(*s)
}

func (s *USet) Add(x int) {
	(*s)[x] = true
}

func (s *USet) Delete(x int) {
	(*s)[x] = false
}

// Find an item in set, returns item if found and true, false if not
// returned item is useful, for example, so one can add (key, value)
// pairs to set and search on key
func (s *USet) Find(x int) (int, bool) {
	return x, (*s)[x] == true
}
