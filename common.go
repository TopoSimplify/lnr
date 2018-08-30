package lnr

import (
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/opts"
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/state"
)

const NullFId = -9

const (
	x = iota
	y
)

type ScoreFn func(coordinates geom.Coords) (int, float64)

type Polygonal interface {
	Coordinates() []*geom.Point
	Polyline() *pln.Polyline
}

type Linegen interface {
	Id() int
	Options() *opts.Opts
	Simple() []int
	State()  *state.State
}

type Linear interface {
	Polygonal
	Linegen
}
