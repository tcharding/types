// Interfaces for Abstract Data Types

type Stack interface {
	Push(x interface{})
	Pop() (interface{}, bool)
	Peek() (interface{}, bool)
	IsEmpty() bool
}

type Queue interface {
	Enqueue(x interface{})
	Dequeue() (interface{}, bool)
}

type Deque interface {
	EnqueueF(x interface{})
	DequeueF() (interface{}, bool)
	EnqueueR(x interface{})
	DequeueR() (interface{}, bool)
}

type List interface {
	Len() int
	Get(i int) interface{}
	Set(i int, x interface{})
	Insert(i int, x interface{})
	Delete(i int)
}

// Unsorted Set
type USet interface {
	Len() int
	Add(x interface{}) bool
	Delete(x interface{})

	// Find an item in set, returns item if found and true, false if not
	// returned item is useful, for example, so one can add (key, value)
	// pairs to set and search on key
	Find(x interface{}) (interface{}, bool)
}

// Sorted Set
type Sset interface {
	Size() int
	Add(x interface{}) bool
	Delete(x interface{})
	Find(x interface{}) (interface{}, error) // returns the successor if not found
}

// Priority Queue
type PriorityQueue interface {
	Size() int
	Insert(x interface{}) bool
	Delete(x interface{})
	Minimum(x interface{}) // this should have time complexity O(1)
}

// Bag: an un-ordered collection of non-unique items
type Bag interface {
	Size() int
	Add(x interface{}) bool
	Delete(x interface{})
	Find(x interface{}) (interface{}, bool)
	FindAll(x interface{}) []interface{}
}
