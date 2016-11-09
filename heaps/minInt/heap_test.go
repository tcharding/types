package heap

import "testing"

func TestExtract(t *testing.T) {
	xs := []int{7, 5, 1, 9, 0, 2, 3, 6, 4, 8}
	h := MakeHeap(xs)
	for i := 0; i < 10; i++ {
		v, ok := h.ExtractMin()
		if !ok {
			t.Errorf("ok error\n")
		}
		if v != i {
			t.Errorf("h.ExtractMin() = %d want: %d\n", v, i)
		}
	}
}
