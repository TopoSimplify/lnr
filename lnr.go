package lnr

import (
    "simplex/pln"
    "simplex/opts"
    "github.com/intdxdt/geom"
    "simplex/nque"
)

type ScoreFn func(coordinates []*geom.Point) (int, float64)

type Polygonal interface {
    Coordinates() []*geom.Point
    Polyline() *pln.Polyline
}

type Linegen interface {
    Id() string
    Options() *opts.Opts
    Simple() []int
}

type NodeQueue interface {
    NodeQueue() *nque.Queue
}

type Linear interface {
    Polygonal
    Linegen
    NodeQueue
}
