package lnr

import (
	"simplex/rng"
	"simplex/pln"
	"github.com/intdxdt/geom"
)

//Linear interface
type Linear interface {
	Coordinates() []*geom.Point
	Polyline() *pln.Polyline
	Score(Linear, *rng.Range) (int, float64)
}
