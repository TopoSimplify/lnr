package lnr

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/mbr"
)

const (
	x = iota
	y
)

const (
	collinear        = iota
	segmentIntersect
	vertexIntersect
)

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

//envelope of segment
func BBox(a, b *geom.Point) *mbr.MBR {
	return mbr.NewMBR(a[x], a[y], b[x], b[y])
}
