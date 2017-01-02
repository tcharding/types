// Weighted Graph, nodes are labelled from 0 to N-1.
package main

import "fmt"

type Graph struct {
	NumNodes int
	Edges    [][]Edge
}

type Edge struct {
	From   int
	To     int
	Weight int
}

// NewGraph: Create graph with n nodes.
func NewGraph(n int) *Graph {
	return &Graph{
		NumNodes: n,
		Edges:    make([][]Edge, n),
	}
}

// AddEdge: Add an edge from u to v.
func (g *Graph) AddEdge(u, v, w int) {
	g.Edges[u] = append(g.Edges[u], Edge{From: u, To: v, Weight: w})

	// For undirected graph add edge from v to u.
	// g.Edges[v] = append(g.Edges[v], Edge{From: v, To: u, Weight: w})
}

func (g *Graph) adjacentEdgesExample() {
	u := 0 // example node label.

	fmt.Printf("Printing all edges adjacent to %d\n", u)
	// Loop over edges adjacent to node u.
	for _, e := range g.Edges[u] {
		fmt.Printf("Edge: %d -> %d (%d)\n", e.From, e.To, e.Weight)
	}
	fmt.Println()

	fmt.Println("Printing all edges.")
	// Loop over all edges.
	for _, adjacent := range g.Edges {
		for _, e := range adjacent {
			fmt.Printf("Edge: %d -> %d (%d)\n", e.From, e.To, e.Weight)
		}
	}
	fmt.Println()
}
