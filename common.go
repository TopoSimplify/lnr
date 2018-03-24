package lnr

import (
	"simplex/pln"
	"simplex/opts"
	"github.com/intdxdt/math"
	"github.com/intdxdt/geom"
)

const NullFId = ""

const (
	x = iota
	y
)

const (
	collinear    = iota
	seginters
	vertexinters
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



//clamp to zero if float is near zero
func snap_to_zero(v float64) float64 {
	if math.FloatEqual(v, 0.0) {
		v = 0.0
	}
	return v
}

//clamp to zero or one
func snap_to_zero_or_one(v float64) float64 {
	if math.FloatEqual(v, 0.0) {
		v = 0.0
	} else if math.FloatEqual(v, 1.0) {
		v = 1.0
	}
	return v
}


