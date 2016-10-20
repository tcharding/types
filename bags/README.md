# Implementations of the Bag interface

**Bag: an un-ordered collection of non-unique items**

type Bag interface {
	Size() int
	Add(x interface{}) bool
	Delete(x interface{})
	Find(x interface{}) (interface{}, bool)
	FindAll(x interface{}) []interface{}
}
    
** see ../interfaces.go **

## integer

bag of integers
