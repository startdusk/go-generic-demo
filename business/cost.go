package business

import (
	"golang.org/x/exp/constraints"
)

func (s Solar) Cost() float64 {
	return s.Netto * 0.4
}

func (w Wind) Cost() float64 {
	return w.Netto * 0.8
}

type Number interface {
	constraints.Float | constraints.Integer
}

// Cost multiplies usage with netto to compute the cost.
func Cost[T Number](usage, netto T) T {
	cost := usage * netto
	return cost
}
