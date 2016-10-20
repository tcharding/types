# Implementations of the USet interface

// Sorted Set
type Sset interface {
	Size() int
	Add(x interface{}) bool
	Delete(x interface{}) bool

	// Find an item in set, returns item if found and true, if not found
	// returns successor item and true. If no successor returns false
	Find(x interface{}) (interface{}, error)
}

** see ../interfaces.go **

## string

ordered set of strings
