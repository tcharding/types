# Implementations of the USet interface

type USet interface {
	Size() int
	Add(x interface{}) bool
	Delete(x interface{}) bool
	Find(x interface{}) (interface{}, bool)
}

** see ../interfaces.go **

## integer

Un-ordered set of integers (i.e golang map type)
