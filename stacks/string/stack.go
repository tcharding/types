// stack of strings
package stack

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new string onto the stack
func (s *Stack) Push(x string) {
	*s = append(*s, x)
}

// Pop: remove and return top element of stack, return false if stack is empty
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	i := len(*s) - 1
	x := (*s)[i]
	*s = (*s)[:i]

	return x, true
}

// Peek: return top element of stack, return false if stack is empty
func (s *Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	i := len(*s) - 1
	x := (*s)[i]

	return x, true
}
