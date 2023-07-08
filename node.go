package graph

import (
	"time"
)

type Node interface {
	Step(dt time.Duration, step uint, args ...float64) float64
}

type BasicNode struct {
	Node
}

func (n BasicNode) Step(_ time.Duration, _ uint, args ...float64) float64 {
	total := 0.
	for _, v := range args {
		total += v
	}
	PrintIf(DEBUG, "outputting %v\n", total)
	return total
}
