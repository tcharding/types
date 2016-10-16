package stack

import "testing"

func TestStack(t *testing.T) {
	var stack Stack

	stack.Push("this")
	stack.Push("and")
	stack.Push("that")

	var wants = []string{"that", "and", "this"}

	if stack.IsEmpty() {
		t.Errorf("stack.IsEmpty returned true for non-empty stack\n")
	}

	for _, want := range wants {
		got, ok := stack.Peek()
		if !ok {
			t.Errorf("Peek returned not ok for non-empty stack\n")
		}
		if got != want {
			t.Errorf("Peek() returned: %v want: %v\n", got, want)
		}

		got, ok = stack.Pop()
		if !ok {
			t.Errorf("Pop returned not ok for non-empty stack\n")
		}
		if got != want {
			t.Errorf("Pop() returned: %v want: %v\n", got, want)
		}
	}

	if !stack.IsEmpty() {
		t.Errorf("stack.IsEmpty returned false for empty stack\n")
	}
}
