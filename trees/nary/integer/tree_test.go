package tree

import (
	"fmt"
	"os"
	"testing"
)

const verbose = false

func Test(t *testing.T) {
	root := NewNode(0)
	one := root.AddChild(1)
	root.AddChild(2)
	root.AddChild(3)
	one.AddChild(4)
	one.AddChild(5)

	want := 6
	if got := root.Size(); got != want {
		t.Errorf("t.Root.Size = %d want: %d\n", got, want)
	}

	want = 3
	if got := one.Size(); got != want {
		t.Errorf("t.Root.Size = %d want: %d\n", got, want)
	}

}

func TestBasic(t *testing.T) {
	root := NewNode(1)
	root.AddChild(2)
	root.AddChild(3)
	want := 3
	if got := root.Size(); got != want {
		t.Errorf("t.Root.Size = %d want: %d\n", got, want)
	}
	if verbose {
		fmt.Fprintf(os.Stderr, "%s\n", root)
	}

}

// func Test(t *testing.T) {
// 	three := newNode(3)
// 	three.AddChild(4)
// 	three.AddChild(5)
// 	three.AddChild(6)
// 	t.Root := newNode(1)
// 	t.Root.AddChild(2)
// 	t.Root.AddChildNode(three)

// 	if verbose {
// 		fmt.Fprintf(os.Stderr, "%s\n", t.Root)
// 	}
// }
