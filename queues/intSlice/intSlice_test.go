package intSlice

import "testing"

func TestQueue(t *testing.T) {
	var q Queue

	val := 1
	q.Enqueue(val)
	x, _ := q.Dequeue()
	if x != val {
		t.Errorf("queue and dequeue of single integer failed\n")
	}

	for i := 0; i < 1000; i++ {
		val := 1
		q.Enqueue(val)
		x, _ := q.Dequeue()
		if x != val {
			t.Errorf("loop queue and dequeue of single integer failed, index: %d\n", i)
		}
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	var wants = []int{1, 2, 3}

	if q.isEmpty() {
		t.Errorf("q.isEmpty returned true for non-empty queue\n")
	}

	for _, want := range wants {
		got, ok := q.Dequeue()
		if !ok {
			t.Errorf("Dequeue returned not ok for non-empty queue\n")
		}
		if got != want {
			t.Errorf("Dequeue() returned: %d want: %d\n", got, want)
		}
	}

	if !q.isEmpty() {
		t.Errorf("q.IsEmpty returned false for empty queue\n")
	}
}
