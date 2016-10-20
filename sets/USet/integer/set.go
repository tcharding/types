// USet: Un-ordered set of Integers, implements the USet interface

// This is just a wrapper for Go's map type

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

func (s *USet) Find(x int) (int, bool) {
	return x, (*s)[x] == true
}
