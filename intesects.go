package lnr

import "github.com/intdxdt/mbr"

//do two lines intersect line segments a && b with
//vertices sa, sb, oa, ob
func Intersects(sa, sb, oa, ob []float64) bool {
	return intersects(sa, sb, oa, ob, false)
}

func intersects(sa, sb, oa, ob []float64, extln bool) bool {
	var bln = false
	var a, b, d,
	x1, y1, x2, y2,
	x3, y3, x4, y4 = segseg_intersect_abdxy(sa, sb, oa, ob )

	//snap to zero if near -0 or 0
	a = snap_to_zero(a)
	b = snap_to_zero(b)
	d = snap_to_zero(d)

	if d == 0 {
		if a == 0.0 && b == 0.0 {
			abox := mbr.NewMBR(x1, y1, x2, y2)
			bbox := mbr.NewMBR(x3, y3, x4, y4)
			bln = abox.Intersects(bbox)
		}
		return bln
	}
	//intersection along the the seg or extended seg
	ua := snap_to_zero_or_one(a / d)
	ub := snap_to_zero_or_one(b / d)

	ua_0_1 := 0.0 <= ua && ua <= 1.0
	ub_0_1 := 0.0 <= ub && ub <= 1.0
	bln = (ua_0_1 && ub_0_1) || extln
	return bln
}
