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

type VBits uint8

const (
	OtherB VBits = 1 << iota // 1 << 0 == 0001
	OtherA                   // 1 << 1 == 0010
	SelfB                    // 1 << 2 == 0100
	SelfA                    // 1 << 3 == 1000
)
const Intersects VBits = 0
const (
	SelfMask  = SelfA | SelfB
	OtherMask = OtherA | OtherB
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
