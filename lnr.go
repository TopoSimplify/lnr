package lnr

import (
	"simplex/rng"
	"simplex/pln"
	"simplex/opts"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/deque"
)

type ScoreFn func(Linear, *rng.Range) (int, float64)

type Polygonal interface {
	Coordinates() []*geom.Point
	Polyline() *pln.Polyline
}

type Linegen interface {
	Id() string
	Score(Linear, *rng.Range) (int, float64)
	Options() *opts.Opts
	Simple() []int
}

type NodeQueue interface {
	NodeQueue() *deque.Deque
}

type Linear interface {
	Polygonal
	Linegen
	NodeQueue
}
