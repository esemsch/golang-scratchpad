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

type Edge struct {
	a Node
	b Node
}

func (n Node) ID() int { return n.id }

func (e Edge) From() graph.Node { return graph.Node(e.a) }

func (e Edge) To() graph.Node { return graph.Node(e.b) }

func (e Edge) Weight() float64 { return math.Sqrt(math.Pow(e.a.x-e.b.x, 2) + math.Pow(e.a.y-e.b.y, 2)) }

func main() {
	n1 := Node{x: 0.0, y: 0.0, id: 1}
	n2 := Node{x: 1.0, y: 1.0, id: 2}
	e := Edge{a: n1, b: n2}

	fmt.Println(e.Weight())
}
