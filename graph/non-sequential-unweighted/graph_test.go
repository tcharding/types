package main

import (
	"fmt"
	"testing"
)

var (
	nNodes  = 10
	verbose = true
)

func Test(t *testing.T) {
	g := NewGraph()

	g.AddEdge(0, 3)
	g.AddEdge(0, 4)
	g.AddEdge(9, 7)
	g.AddEdge(1, 2)

	if verbose {
		fmt.Printf("Nodes: [")
		for id := range g.Nodes {
			fmt.Printf("%d ", id)
		}
		fmt.Printf("]\n")

		g.adjacentEdgesExample()
	}
}
