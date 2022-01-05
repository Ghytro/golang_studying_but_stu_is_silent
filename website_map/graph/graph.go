package graph

import (
	"fmt"
	"strings"
)

type Edge struct {
	Source string
	Target string
}

type Graph struct {
	edgesSet map[Edge]int // the int value is the amount of the edges
}

func NewGraph() *Graph {
	result := new(Graph)
	result.edgesSet = make(map[Edge]int)
	return result
}

func (g Graph) GetEdges() []Edge {
	var result []Edge
	for k, v := range g.edgesSet {
		for i := 0; i < v; i++ {
			result = append(result, k)
		}
	}
	return result
}

func (g Graph) AddEdge(e Edge) {
	g.edgesSet[e]++
}

func (g Graph) ToGraphviz(graphName string) string {
	var sb strings.Builder
	const indent = "    "
	sb.WriteString("digraph " + graphName + " {\n")
	for _, e := range g.GetEdges() {
		sb.WriteString(fmt.Sprintf("%s\"%s\" -> \"%s\"\n", indent, e.Source, e.Target))
	}
	sb.WriteString("}")
	return sb.String()
}
