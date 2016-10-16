package stack

import "testing"

func TestStack(t *testing.T) {
	var stack Stack

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	var wants = []int{3, 2, 1}

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
