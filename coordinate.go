package main

import (
	"math"
)

type Coordinate struct {
	X float64
	Y float64
}

func distance(c1 Coordinate, c2 Coordinate) float64 {
	return math.Sqrt((c1.X-c2.X)*(c1.X-c2.X) + (c1.Y-c2.Y)*(c1.Y-c2.Y))
}
