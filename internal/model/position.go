package model

import "math"

type Position struct {
	X, Y float64
}

func (p Position) DistanceTo(position Position) float64 {
	x := math.Abs(p.X - position.X)
	y := math.Abs(p.Y - position.Y)
	return math.Sqrt(x*x + y*y)
}
