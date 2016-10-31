# nodes

Nodes used for building trees.

e.g  

```go
type Node struct {
    key int
    parent *Node
    left   *Node
    right  *Node
}
```

While being simple to implement, and typical in literature, trees have a
couple of disadvantages when implemented thus in Golang. Firstly the zero value
is not useful (usually) since it is a node with key value 0. Secondly, although
only a minor disadvantage, functions that modify the node or any node below this
node usually must return the node so that changes can be stored

e.g

```go
    var n Node <!-- key = 0 -->
    n = n.Add(1)
```

Both these shortcomings can be overcome by wrapping a node in another struct,
e.g

```go
    type Tree struct {
        root *Node
    }
```

<!-- See [here]() for example code using this design-->

