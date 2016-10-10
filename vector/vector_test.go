package vector

import "testing"

func TestString(t *testing.T) {
	var v Vector = []string{"this", "one"}
	got, _ := v.String()
	want := "[\"this\", \"one\"]"
	if got != want {
		t.Errorf("v.String; got: %s want: %s\n", got, want)
	}
}

func TestEqual(t *testing.T) {
	var v Vector = []string{"this", "one"}
	var w Vector = []string{"this", "one"}
	var x Vector = []string{"this", "one", "not", "this", "one"}

	var tests = []struct {
		a    Vector
		b    Vector
		want bool
	}{
		{v, v, true},
		{v, w, true},
		{v, x, false},
	}
	for _, test := range tests {
		a := test.a
		b := test.b
		want := test.want
		if a.Equal(b) != want {
			t.Errorf("a.Equal(b) failed a: %s b: %s want: %v\n", a, b, want)
		}
	}

}
