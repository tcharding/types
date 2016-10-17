// intSlice
//
// implements Queue interface
//
// type Queue interface {
// 	Enqueue(x interface{})
// 	Dequeue() (interface{}, bool)
// }

package intSlice

type Queue []int

func (q *Queue) Enqueue(x int) {
	*q = append(*q, x)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.isEmpty() {
		return 0, false
	}

	x := (*q)[0]
	*q = (*q)[1:]
	return x, true
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}
