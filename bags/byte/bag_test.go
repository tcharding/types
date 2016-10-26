package bag

import "testing"

func TestBagUnion(t *testing.T) {
	var tests = []struct {
		s    string
		t    string
		want string
	}{
		{"a", "b", "ab"},
		{"aa", "bb", "aabb"},
		{"abc", "a", "abca"},
	}
	for _, test := range tests {
		bagS := makeBag(test.s)
		bagT := makeBag(test.t)
		want := makeBag(test.want)
		if got := bagS.union(bagT); !bagEq(got, want) {
			t.Errorf("union(%s, %s) = %v want %s\n",
				test.s, test.t, got, test.want)
		}
	}
}

func TestBagIntersection(t *testing.T) {
	var tests = []struct {
		s    string
		t    string
		want string
	}{
		{"a", "b", ""},
		{"ac", "bc", "c"},
		{"aab", "abc", "ab"},
	}
	for _, test := range tests {
		bagS := makeBag(test.s)
		bagT := makeBag(test.t)
		want := makeBag(test.want)
		if got := bagS.intersection(bagT); !bagEq(got, want) {
			t.Errorf("intersection(%s, %s) = %v want %s\n",
				test.s, test.t, got, test.want)
		}
	}
}

func TestBagDifference(t *testing.T) {
	var tests = []struct {
		s    string
		t    string
		want string
	}{
		{"a", "b", "a"},
		{"ac", "bc", "a"},
		{"aab", "abc", "a"},
	}
	for _, test := range tests {
		bagS := makeBag(test.s)
		bagT := makeBag(test.t)
		want := makeBag(test.want)
		if got := bagS.difference(bagT); !bagEq(got, want) {
			t.Errorf("difference(%s, %s) = %v want %s\n",
				test.s, test.t, got, test.want)
		}
	}
}
