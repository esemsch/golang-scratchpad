package main

import (
	"fmt"
	"github.com/gonum/graph"
	"github.com/gonum/graph/path"
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

func euclDist(a, b Node) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}
func (e Edge) Weight() float64 { return euclDist(e.a, e.b) }

type Graph struct {
	nodes []graph.Node
	edges map[graph.Node][]Edge
}

func (g Graph) Has(node graph.Node) bool {
	_, ok := g.edges[node]
	return ok
}

func (g Graph) Nodes() []graph.Node { return g.nodes }

func (g Graph) From(node graph.Node) []graph.Node {
	es, ok := g.edges[node]
	retVal := []graph.Node{}
	if ok {
		for _, e := range es {
			retVal = append(retVal, e.b)
		}
	}
	return retVal
}

func (g Graph) HasEdgeBetween(x, y graph.Node) bool { return g.Edge(x, y) != nil }

func (g Graph) Edge(x, y graph.Node) graph.Edge {
	ex, okx := g.edges[x]
	ey, oky := g.edges[y]
	if okx {
		for _, e := range ex {
			if e.b == y {
				return e
			}
		}
	}
	if oky {
		for _, e := range ey {
			if e.b == x {
				return e
			}
		}
	}
	return nil

}

func (g Graph) Weight(x, y graph.Node) (w float64, ok bool) {
	if x.ID() == y.ID() {
		return 0.0, true
	}
	edge := g.Edge(x, y)
	if edge != nil {
		return edge.Weight(), true
	} else {
		return 0.0, false
	}
}

func main() {
	nodes := []Node{}
	edges := make(map[graph.Node][]Edge)

	const X_SIZE = 1000
	const Y_SIZE = 1000
	for i := 0; i < X_SIZE; i++ {
		for j := 0; j < Y_SIZE; j++ {
			nodes = append(nodes, Node{x: float64(i), y: float64(j), id: i*X_SIZE + j})
		}
	}

	edgeCounter := 0
	for i := 0; i < X_SIZE; i++ {
		for j := 0; j < Y_SIZE; j++ {
			node := nodes[i*X_SIZE+j]
			es := generateEdges(i, j, nodes, X_SIZE, Y_SIZE)
			edges[node] = es
			edgeCounter += len(es)
		}
	}

	graphNodes := make([]graph.Node, len(nodes))
	for i, n := range nodes {
		graphNodes[i] = n
	}

	fmt.Printf("%d nodes, %d edges\n", len(nodes), edgeCounter)

	G := Graph{nodes: graphNodes, edges: edges}

	from := nodes[0]
	to := nodes[X_SIZE*Y_SIZE-2]
	//to := nodes[X_SIZE*(Y_SIZE-1)+Y_SIZE/2]
	heuristic := func(n, m graph.Node) float64 { return euclDist(n.(Node), m.(Node)) }
	//heuristic := func(n, m graph.Node) float64 { return 0 }
	path, exp := path.AStar(from, to, G, heuristic)

	nodesOnPath, dist := path.To(to)
	fmt.Println(nodesOnPath, "Distance:", dist, "Nodes expanded:", exp)
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
