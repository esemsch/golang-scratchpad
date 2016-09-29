package main

import (
	"fmt"
	"github.com/gonum/graph"
	"math"
)

type Node struct {
	x  float64
	y  float64
	id int
}

func (n Node) ID() int { return n.id }

type Edge struct {
	a Node
	b Node
}

func (e Edge) From() graph.Node { return graph.Node(e.a) }

func (e Edge) To() graph.Node { return graph.Node(e.b) }

func (e Edge) Weight() float64 { return math.Sqrt(math.Pow(e.a.x-e.b.x, 2) + math.Pow(e.a.y-e.b.y, 2)) }

type Graph struct {
	nodes []Node
	edges map[Node][]Edge
}
											// Edge returns the edge from u to v if such an edge
												// exists and nil otherwise. The node v must be directly
													// reachable from u as defined by the From method.
													Edge(u, v Node) Edge
func (g Graph) Has(node Node) bool {
	_,ok := edges[node]
	return ok
}

func (g Graph) Nodes() []Node { return nodes }

func (g Graph) From(node Node) []Node {
	es,ok := g.edges[node]
	retVal := []Node{}
	if ok {
		for _,e := range es {
			retVal = append(retVal,e.b)
		}
	}
	return retVal
}

func (g Graph) HasEdgeBetween(x, y Node) bool {return Edge(x,y)!=nil}

func (g Graph) Edge(u, v Node) Edge {
	ex := g.From(x)
	ey := g.From(y)
	for e := range ex {
		if e.b == y {
			return e
		}
	}
	for e := range ey {
		if e.b == x {
			return e
		}
	}
	return nil

}

func main() {
	nodes := []Node{}
	edges := []Edge{}

	const X_SIZE = 2
	const Y_SIZE = 2
	for i := 0; i < X_SIZE; i++ {
		for j := 0; j < Y_SIZE; j++ {
			nodes = append(nodes, Node{x: float64(i), y: float64(j), id: i*X_SIZE + j})
		}
	}

	for i := 0; i < X_SIZE; i++ {
		for j := 0; j < Y_SIZE; j++ {
			edges = append(edges, generateEdges(i, j, nodes, X_SIZE, Y_SIZE)...)
		}
	}

	fmt.Printf("%d nodes, %s edges\n", len(nodes), len(edges))


}

func generateEdges(x, y int, nodes []Node, X_SIZE, Y_SIZE int) []Edge {
	node := nodes[x*X_SIZE+y]
	edges := []Edge{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			nx := x + i
			ny := y + j
			if nx >= 0 && nx < X_SIZE && ny >= 0 && ny < Y_SIZE {
				node2 := nodes[nx*X_SIZE+ny]
				if node != node2 {
					edges = append(edges, Edge{a: node, b: node2})
				}
			}
		}
	}
	return edges
}
