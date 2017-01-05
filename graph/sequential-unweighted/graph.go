// Unweighted Graph, nodes are labelled from 0 to N-1.
package main

import "fmt"

type Graph struct {
	NumNodes int
	Edges    [][]int
}

// NewGraph: Create graph with n nodes.
func NewGraph(n int) *Graph {
	return &Graph{
		NumNodes: n,
		Edges:    make([][]int, n),
	}
}

// AddEdge: Add an edge from u to v.
func (g *Graph) AddEdge(u, v int) {
	g.Edges[u] = append(g.Edges[u], v)

	// For undirected graph add edge from v to u.
	// g.Edges[v] = append(g.Edges[v], u)
}

func (g *Graph) adjacentEdgesExample() {
	u := 0 // Example node label.

	fmt.Printf("Printing all edges adjacent to Node: %d\n", u)
	for _, v := range g.Edges[u] {
		// Edge exists from u to v.
		fmt.Printf("Edge: %d -> %d\n", u, v)
	}

	fmt.Println("Printing all edges in graph.")
	for u, adjacent := range g.Edges { // Nodes are labelled 0 to N-1.
		for _, v := range adjacent {
			// Edge exists from u to v.
			fmt.Printf("Edge: %d -> %d\n", u, v)
		}
	}
}
