# Implementations of the list interface

type List interface {
	Len() int
	Get(i int) interface{}
	Set(i int, x interface{})
	Insert(i int, x interface{})
	Delete(i int)
}

** see ../interfaces.go **

## initial

initial attempt, implement functionality without concern for efficiency. 
