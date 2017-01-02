// Unweighted Graph, nodes are labelled with arbitrary integer ID's.
package main

import "fmt"

type ID int

type Graph struct {
	Nodes map[ID]struct{}
	Edges map[ID]map[ID]struct{}
}

// NewGraph: Create graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[ID]struct{}),
		Edges: make(map[ID]map[ID]struct{}),
	}
}

// AddNode: Add node id to graph, return true if added (ID's are unique).
func (g *Graph) AddNode(n ID) bool {
	if _, ok := g.Nodes[n]; ok {
		return false
	}
	g.Nodes[n] = struct{}{}
	return true
}

// AddEdge: Add an edge from u to v.
func (g *Graph) AddEdge(u, v ID) {
	if _, ok := g.Nodes[u]; !ok {
		g.AddNode(u)
	}
	if _, ok := g.Nodes[v]; !ok {
		g.AddNode(v)
	}

	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = make(map[ID]struct{})
	}
	g.Edges[u][v] = struct{}{}

	// For undirected graph add edge from v to u.
	// if _, ok := g.Edges[v]; !ok {
	// 	g.Edges[v] = make(map[ID]struct{})
	// }
	// g.Edges[v][u] = struct{}{}
}

func (g *Graph) adjacentEdgesExample() {
	u := ID(0) // example node ID.

	// Loop over edges adjacent to node u.
	fmt.Printf("Printing all edges adjacent to %d\n", u)
	for v := range g.Edges[u] {
		// Edge exists from u to v.
		fmt.Printf("Edge: %d -> %d\n", u, v)
	}
	fmt.Println()

	// Loop over all edges.
	fmt.Println("Printing all edges.")
	for u, m := range g.Edges { // Nodes are labelled 0 to N-1.
		for v := range m {
			// Edge exists from u to v.
			fmt.Printf("Edge: %d -> %d\n", u, v)
		}
	}
	fmt.Println()
}
