# USets

## Implementations of the USet interface

type USet interface {
	Len() int
	Add(x interface{}) bool
	Delete(x interface{})
	Find(x interface{}) (interface{}, bool)
}

** see ../interfaces.go **

### initial

initial attempt, implement functionality without concern for efficiency. 
