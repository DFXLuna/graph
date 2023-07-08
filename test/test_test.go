package test

import (
	"egrant/graph"
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	layer0 := make([]graph.Node, 2)
	layer0[0] = graph.BasicNode{}
	layer0[1] = graph.BasicNode{}

	nodes := [][]graph.Node{layer0, layer0, layer0}
	weights := make([][][]float64, 3)
	for i := 0; i < 3; i++ {
		node0 := []float64{1, .5}
		node1 := []float64{2, 3}
		layer := [][]float64{node0, node1}
		weights[i] = layer
	}
	g := graph.NewGraph(nodes, weights)
	fmt.Printf("outs: %v\n", g.Step(1, 1, 2))
}
