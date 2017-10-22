package lnr

import (
	"simplex/rng"
	"simplex/pln"
	"simplex/opts"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/sset"
	"github.com/intdxdt/deque"
)

//score function
type ScoreFn func(Linear, *rng.Range) (int, float64)

type NodeQueue interface {
	NodeQueue() *deque.Deque
}

type Polygonal interface {
	Coordinates() []*geom.Point
	Polyline() *pln.Polyline
}

type Lingen interface {
	Score(Linear, *rng.Range) (int, float64)
	Options() *opts.Opts
}

//Linear interface
type Linear interface {
	NodeQueue
	Polygonal
	Lingen
}

type SimpleAlgorithm interface {
	Simple() *sset.SSet
}
