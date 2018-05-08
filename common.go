package lnr

import (
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/opts"
	"github.com/intdxdt/geom"
)

const NullFId = ""

const (
	x = iota
	y
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

type Linear interface {
	Polygonal
	Linegen
}
