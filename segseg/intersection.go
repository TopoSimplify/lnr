package lnr

import (
	"fmt"
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom"
	"sort"
)

//do two lines intersect line segments a && b with
//vertices lna0, lna1 and lnb0, lnb1
func SegIntersection(sa, sb, oa, ob *geom.Point) []*IntPt {
	var coords = make([]*IntPt, 0)
	var a, b, d = segseg_abd(sa[:], sb[:], oa[:], ob[:])

	//snap to zero if near -0 or 0
	a = snap_to_zero(a)
	b = snap_to_zero(b)
	d = snap_to_zero(d)

	// Are the line coincident?
	if d == 0 {
		if a == 0 && b == 0 {
			var abox = BBox(sa, sb)
			var bbox = BBox(oa, ob)
			if region, ok := abox.Intersection(bbox); ok {
				update_coords_inbounds(bbox, sa, &coords, region, "sa-.-.-.")
				update_coords_inbounds(bbox, sb, &coords, region, ".-sb-.-.")
				update_coords_inbounds(abox, oa, &coords, region, ".-.-oa-.")
				update_coords_inbounds(abox, ob, &coords, region, ".-.-.-ob")
			}
		}
		sort.Sort(IntPts(coords))
		return coords
	}
	// is the intersection along the the segments
	var ua = snap_to_zero_or_one(a / d)
	var ub = snap_to_zero_or_one(b / d)

	var ua_0_1 = 0.0 <= ua && ua <= 1.0
	var ub_0_1 = 0.0 <= ub && ub <= 1.0

	if ua_0_1 && ub_0_1 {
		var postfix = inter_postfix(ua, ub)
		var prefix = "x"
		// intersection point is within range of lna && lnb ||  by extension
		if (ua == 0 || ua == 1) || (ub == 0 || ub == 1) {
			prefix = "v"
		}
		var pt = &IntPt{
			geom.NewPointXY(
				sa[x]+ua*(sb[x]-sa[x]),
				sa[y]+ua*(sb[y]-sa[y]),
			),
			prefix + "-" + postfix,
		}
		coords = append(coords, pt)
	}
	sort.Sort(IntPts(coords))
	return coords
}

func segseg_abd(sa, sb, oa, ob []float64) (float64, float64, float64) {

	var x1, y1, x2, y2, x3, y3, x4, y4, d, a, b float64

	x1, y1 = sa[x], sa[y]
	x2, y2 = sb[x], sb[y]

	x3, y3 = oa[x], oa[y]
	x4, y4 = ob[x], ob[y]

	d = ((y4 - y3) * (x2 - x1)) - ((x4 - x3) * (y2 - y1))
	a = ((x4 - x3) * (y1 - y3)) - ((y4 - y3) * (x1 - x3))
	b = ((x2 - x1) * (y1 - y3)) - ((y2 - y1) * (x1 - x3))

	return a, b, d
}

func inter_postfix(ua, ub float64) string {
	var sa, sb, oa, ob = ".", ".", ".", "."
	if ua == 0 {
		sa = "sa"
	} else if ua == 1 {
		sb = "sb"
	}

	if ub == 0 {
		oa = "oa"
	} else if ub == 1 {
		ob = "ob"
	}
	return fmt.Sprintf("%v-%v-%v-%v", sa, sb, oa, ob)
}

//updates coords that are in bounds
func update_coords_inbounds(bounds *mbr.MBR, point *geom.Point, intpts *[]*IntPt, region *mbr.MBR, postfix string) {
	if bounds.ContainsXY(point[x], point[y]) {
		var prefix = "c"
		if region.IsPoint() {
			prefix = "v"
		}
		pnt := &IntPt{point.Clone(), prefix + "-" + postfix}
		*intpts = append(*intpts, pnt)
	}
}
