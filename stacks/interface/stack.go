package stack

// All types satisfy the empty interface, so we can store anything here.
type Stack []interface{}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

// Pop: remove and return top element of stack, return false if stack is empty
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		var empty interface{}
		return empty, false
	}

	i := len(*s) - 1
	x := (*s)[i]
	*s = (*s)[:i]

	return x, true
}

// Peek: return top element of stack, return false if stack is empty
func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		var empty interface{}
		return empty, false
	}

	i := len(*s) - 1
	x := (*s)[i]

	return x, true
}
