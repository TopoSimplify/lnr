package lnr

import (
	"github.com/intdxdt/sset"
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/math"
	"github.com/intdxdt/geom"
)

const (
	collinear    = iota
	seginters
	vertexinters
)

type interPt struct {
	*geom.Point
	inter int
}

func (p *interPt) isCollinear() bool {
	return p.inter == collinear
}

func (p *interPt) isSegmentIntersection() bool {
	return p.inter == seginters
}

func (p *interPt) isVertexIntersection() bool {
	return p.inter == vertexinters
}

//do two lines intersect line segments a && b with
//vertices lna0, lna1 and lnb0, lnb1
func SegIntersection(self, other *geom.Segment) []*interPt {
	var sa, sb, oa, ob = self.A[:], self.B[:], other.A[:], other.B[:]
	var set = sset.NewSSet(lexsort2d)
	var coords = make([]*interPt, 0)
	var a, b, d, x1, y1, x2, y2,
	x3, y3, x4, y4 = segseg_intersect_abdxy(sa, sb, oa, ob)

	//snap to zero if near -0 or 0
	a = snap_to_zero(a)
	b = snap_to_zero(b)
	d = snap_to_zero(d)

	// Are the line coincident?
	if d == 0 {
		if a == 0 && b == 0 {
			var abox = self.BBox()
			var bbox = other.BBox()
			var region, _ = self.BBox().Intersection(other.BBox())
			if abox.Intersects(bbox) {
				update_coords_inbounds(abox, x3, y3, x4, y4, set, region)
				update_coords_inbounds(bbox, x1, y1, x2, y2, set, region)
			}
		}
		set.ForEach(func(o interface{}, _ int) bool {
			coords = append(coords, o.(*interPt))
			return true
		})
		return coords
	}
	// is the intersection along the the segments
	ua := snap_to_zero_or_one(a / d)
	ub := snap_to_zero_or_one(b / d)

	ua_0_1 := 0.0 <= ua && ua <= 1.0
	ub_0_1 := 0.0 <= ub && ub <= 1.0

	if ua_0_1 && ub_0_1 {
		var pt = &interPt{geom.NewPointXY(x1+ua*(x2-x1), y1+ua*(y2-y1)), seginters}
		// intersection point is within range of lna && lnb ||  by extension
		if (ua == 0 || ua == 1) && (ub == 0 || ub == 1) {
			pt.inter = vertexinters
		}

		coords = append(coords, pt)
	}

	return coords
}

//sort 2d coordinates lexicographically
func lexsort2d(apt, bpt interface{}) int {
	var a, b = apt.(*interPt).Point[:], bpt.(*interPt).Point[:]
	var d = a[0] - b[0]
	if math.FloatEqual(d, 0) {
		d = a[1] - b[1]
	} else {
		return lexval(d)
	}

	if math.FloatEqual(d, 0) {
		return 0
	}
	return lexval(d)
}

func lexval(d float64) int {
	if d < 0 {
		return -1
	}
	return 1
}

func segseg_intersect_abdxy(sa, sb, oa, ob []float64) (float64, float64, float64,
	float64, float64, float64, float64,
	float64, float64, float64, float64) {

	var x1, y1, x2, y2, x3, y3, x4, y4, d, a, b float64

	x1, y1 = sa[x], sa[y]
	x2, y2 = sb[x], sb[y]

	x3, y3 = oa[x], oa[y]
	x4, y4 = ob[x], ob[y]

	d = ((y4 - y3) * (x2 - x1)) - ((x4 - x3) * (y2 - y1))
	a = ((x4 - x3) * (y1 - y3)) - ((y4 - y3) * (x1 - x3))
	b = ((x2 - x1) * (y1 - y3)) - ((y2 - y1) * (x1 - x3))
	return a, b, d, x1, y1, x2, y2, x3, y3, x4, y4
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

//updates coords that are in bounds
func update_coords_inbounds(bounds *mbr.MBR, x1, y1, x2, y2 float64, set *sset.SSet, region *mbr.MBR) {
	var pt *interPt
	if bounds.ContainsXY(x1, y1) {
		pt = &interPt{geom.NewPointXY(x1, y1), collinear}
		if region.IsPoint() {
			pt.inter = vertexinters
		}
		set.Add(pt)
	}

	if bounds.ContainsXY(x2, y2) {
		pt = &interPt{geom.NewPointXY(x2, y2), collinear}
		if region.IsPoint() {
			pt.inter = vertexinters
		}
		set.Add(pt)
	}
}
