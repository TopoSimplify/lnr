package lnr

import "github.com/intdxdt/math"

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
