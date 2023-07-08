package graph

import (
	"time"

	gofunctional "github.com/DFXLuna/go-functional"
)

const (
	DEBUG = false
)

// Constraints
// - Graphs are fully connected, feed forward and acyclic
// - Nodes can only connect to nodes in the immediate following layer
type Graph struct {
	Nodes            [][]Node
	Weights          [][][]float64 // [x][y][z] where x is layer, y is node index and z is node index in following layer
	StepNumber       uint
	LargestLayerSize int
}

func NewGraph(nodes [][]Node, weights [][][]float64) *Graph {
	maxSize := func(currMax int, ns []Node) int {
		if len(ns) > currMax {
			return len(ns)
		}
		return 0
	}

	return &Graph{Nodes: nodes, Weights: weights, LargestLayerSize: gofunctional.Foldr(0, maxSize, nodes)}
}

func (g *Graph) Step(dt time.Duration, args ...float64) []float64 {
	g.StepNumber += 1
	prevLayerOuts := make([]float64, g.LargestLayerSize)
	currLayerOuts := make([]float64, g.LargestLayerSize)

	// Do first layer separately to fill prevLayerOuts
	// Layer 0 processing
	for i, n := range g.Nodes[0] {
		prevLayerOuts[i] = n.Step(dt, g.StepNumber, args[i])
	}
	for i := 1; i < len(g.Nodes); i++ { // Process each layer
		PrintIf(DEBUG, "layer %d\n", i)
		for j, n := range g.Nodes[i] { // Process each node
			weightedIns := make([]float64, g.LargestLayerSize)
			for k, v := range prevLayerOuts { // Get weighted outputs from previous layer to use as inputs
				weightedIns[k] = v * g.Weights[i][k][j]
				PrintIf(DEBUG, "Weight[%d][%d][%d] is %v\nWeightedIn is %v\n", i, k, j, g.Weights[i][k][j], weightedIns[k])
			}
			PrintIf(DEBUG, "Stepping node[%d][%d]\n", i, j)
			PrintIf(DEBUG, "weighted ins are %v\n", weightedIns)
			currLayerOuts[j] = n.Step(dt, g.StepNumber, weightedIns...)
		}
		PrintIf(DEBUG, "current outs are %v\n\n----------------\n", currLayerOuts)
		copy(prevLayerOuts, currLayerOuts)
	}
	return prevLayerOuts
}
