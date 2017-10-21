package lnr

import (
	"simplex/rng"
	"simplex/pln"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/sset"
	"simplex/opts"
)

//score function
type ScoreFn func(Linear, *rng.Range) (int, float64)

//Linear interface
type Linear interface {
	Coordinates() []*geom.Point
	Polyline() *pln.Polyline
	Score(Linear, *rng.Range) (int, float64)
	Options() *opts.Opts
}

type SimpleAlgorithm interface {
	Simple() *sset.SSet
}
