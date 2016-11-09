package deque

import "testing"

func TestBasic(t *testing.T) {
	check := func(msg string, got, want int) {
		if got != want {
			t.Errorf("%s: got: %d want: %d\n", msg, got, want)
		}
	}

	var d Deque
	d.EnqueueF(1)
	d.EnqueueB(2)

	if x, ok := d.DequeueF(); ok {
		check("dequeueF", x, 1)
	} else {
		t.Errorf("ok error on non-empty deque\n")
	}
	if x, ok := d.DequeueF(); ok {
		check("dequeueF", x, 2)
	} else {
		t.Errorf("ok error on non-empty deque\n")
	}

	if _, ok := d.DequeueF(); ok {
		t.Errorf("ok error on empty deque\n")
	}
}

func TestBalance(t *testing.T) {
	deque := func(d *Deque, i int) {
		x, _ := d.DequeueF()
		if x != i {
			t.Errorf("deque error: i: %d got: %d\n", i, x)
		}
	}

	ntests := 10
	var d Deque
	for i := 0; i < ntests; i++ {
		// no balance required
		d.EnqueueF(i)
		d.EnqueueB(i)
	}
	// force maximum re-balance
	for i := ntests - 1; i >= 0; i-- {
		deque(&d, i)
	}
	for i := 0; i < ntests; i++ {
		deque(&d, i)
	}

}
