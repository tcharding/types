package list

import "testing"

func Test(t *testing.T) {
	var ls List
	nVals := 5

	// test Insert
	for i := 0; i < nVals; i++ {
		ls.Insert(i, i)
	}

	// test Len
	wantLength := nVals
	if got := ls.Len(); got != wantLength {
		t.Errorf("ls.Len() = %d want: %d\n", got, wantLength)
	}

	// test Get
	for i := 0; i < nVals; i++ {
		x := ls.Get(i)
		if x != i {
			t.Errorf("ls.Get() = %d want: %d\n", x, i)
		}
	}

	// test Set
	for i := 0; i < nVals; i++ {
		want := i * 2
		ls.Set(i, want)
		got := ls.Get(i)
		if got != want {
			t.Errorf("ls.Set(%d), ls.Get() = %d want: %d\n", want, got, want)
		}
	}
}

func TestDelete(t *testing.T) {
	var ls List
	test := func(d int, want []int, msg string) {
		ls.Delete(d)
		if !equal(ls, want) {
			t.Errorf("\nList not correct after ls.Delete(%d): %s\n", d, msg)
		}
		wantLength := len(want)
		if got := ls.Len(); got != wantLength {
			t.Errorf("\nls.Len() not correct after ls.Delete(%d) got: %d want: %d: %s\n",
				d, got, wantLength, msg)
		}
	}

	for i := 0; i < 5; i++ {
		ls.Insert(i, i)
	}

	// List: 0, 1, 2, 3, 4
	test(1, []int{0, 2, 3, 4}, "delete from middle")
	test(2, []int{0, 2, 4}, "delete from middle again")
	test(0, []int{2, 4}, "delete from front")

	test(1, []int{2}, "delete from back")
	test(0, []int{}, "delete last item")
}

func equal(ls List, xs []int) bool {
	for i := range xs {
		if ls.Get(i) != xs[i] {
			return false
		}
	}
	return true
}
