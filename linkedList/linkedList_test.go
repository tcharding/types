package linkedList

import "testing"

func TestAddHead(t *testing.T) {
	var lk LinkedList
	var tests = []struct {
		input string
	}{
		{"this"},
		{"one"},
		{"and"},
		{"that"},
		{"two"},
	}
	counter := 0
	tail := "this"
	for _, test := range tests {
		lk.AddHead(test.input)
		counter++
		if lk.Length() != counter {
			t.Errorf("Length mismatch, got: %d want: %d\n", lk.Length(), counter)
		}
		s := lk.head.value.(string)
		if s != test.input {
			t.Errorf("head mismatch, got: %v want: %v\n", s, test.input)
		}
		s = lk.tail.value.(string)
		if s != tail {
			t.Errorf("tail mismatch, got: %v want: %v\n", s, tail)
		}
	}
}

func TestAddTail(t *testing.T) {
	var lk LinkedList
	var tests = []struct {
		input string
	}{
		{"this"},
		{"one"},
		{"and"},
		{"that"},
		{"two"},
	}
	head := "this"
	counter := 0
	for _, test := range tests {
		lk.AddTail(test.input)
		counter++
		if lk.Length() != counter {
			t.Errorf("Length mismatch, got: %d want: %d\n", lk.Length(), counter)
		}
		s := lk.tail.value.(string)
		if s != test.input {
			t.Errorf("tail mismatch, got: %v want: %v\n", s, test.input)
		}
		s = lk.head.value.(string)
		if s != head {
			t.Errorf("head mismatch, got: %v want: %v\n", s, head)
		}
	}
}

func TestForEach(t *testing.T) {
	var lk LinkedList

	input := []string{"this", "one", "and", "that", "one"}

	addHead(&lk, input...)
	counter := lk.Length() - 1
	fn := func(n *Node) {
		got := n.value.(string)
		want := input[counter]
		if got != want {
			t.Errorf("ForEach: value mismatch, got: %v want: %v\n", got, want)
		}
		counter--
	}
	lk.ForEach(fn)
}

func TestDelete(t *testing.T) {
	var lk LinkedList
	input := []string{"this", "one", "and", "that", "one"}
	addTail(&lk, input...)
	for i, _ := range input {
		lk.Delete(lk.head)
		got := lk.Length()
		want := len(input) - i - 1
		if got != want {
			t.Errorf("length mismatch, got: %v want: %v\n", got, want)
		}
	}
}

// equalFn: returns a function that can be parsed to lk.Delete
func equalFn(s string) func(n *Node) bool {
	fn := func(n *Node) bool {
		if n.value.(string) == s {
			return true
		}
		return false
	}
	return fn
}

func addHead(lk *LinkedList, strings ...string) {
	for _, s := range strings {
		lk.AddHead(s)
	}
}

func addTail(lk *LinkedList, strings ...string) {
	for _, s := range strings {
		lk.AddTail(s)
	}
}
