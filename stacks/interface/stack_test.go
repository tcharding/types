package stack

import "testing"

var stack Stack

func Test(t *testing.T) {
	stack.Push("this")
	stack.Push("and")
	stack.Push("that")

	var tests = []struct {
		want string
	}{
		{"that"},
		{"and"},
		{"this"},
	}
	for _, test := range tests {
		if got := stack.Pop(); got.(string) != test.want {
			t.Errorf("pop() returned: %v want: %v\n", got, test.want)
		}
	}

}
