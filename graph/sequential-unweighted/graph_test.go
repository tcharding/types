package main

import "testing"

var (
	nNodes  = 10
	verbose = true
)

func Test(t *testing.T) {
	g := NewGraph(nNodes)

	g.AddEdge(0, 3)
	g.AddEdge(0, 4)
	g.AddEdge(9, 7)
	g.AddEdge(1, 2)

	if verbose {
		g.adjacentEdgesExample()
	}
}
