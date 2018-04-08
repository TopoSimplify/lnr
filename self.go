package lnr

import (
	"simplex/ctx"
	"simplex/seg"
	"simplex/pln"
	"github.com/intdxdt/iter"
	"github.com/intdxdt/rtree"
)

//Planar and non-planar intersections
func SelfIntersection(polyline *pln.Polyline, planar, nonPlanar bool) *ctx.ContextGeometries {
	var inters = ctx.NewContexts()
	if planar {
		inters.Extend(planarIntersects(polyline).DataView())
	}
	if nonPlanar {
		inters.Extend(nonPlanarIntersection(polyline).DataView())
	}
	return inters
}

//Planar self-intersection
func planarIntersects(polyline *pln.Polyline) *ctx.ContextGeometries {
	var points = make([]*vertex, 0, len(polyline.Coordinates))
	for i, pt := range polyline.Coordinates {
		points = append(points, &vertex{pt, i, NullFId})
	}
	vertices(points).Sort() //O(nlogn)
	var d = 0
	var a, b *vertex
	var indices []int
	var results = ctx.NewContexts()

	var bln bool
	for i, n := 0, len(points); i < n-1; i++ { //O(n)
		a, b = points[i], points[i+1]
		bln = a.Equals2D(b.Point)
		if bln {
			if d == 0 {
				indices = append(indices, a.index, b.index)
				d += 2
			} else {
				indices = append(indices, b.index)
				d += 1
			}
			continue
		}

		if d > 1 {
			var cg = ctx.New(points[i].Point.Clone(), 0, -1).AsPlanarVertex()
			cg.Meta.Planar = iter.SortedIntsSet(indices)
			results.Push(cg)
		}
		d = 0
		indices = indices[:0]
	}
	return results
}

func nonPlanarIntersection(polyline *pln.Polyline) *ctx.ContextGeometries {
	var cache = make(map[[4]int]bool)
	var tree, data = segmentDB(polyline)
	var results = ctx.NewContexts()

	for _, d := range data {
		var s = d.(*seg.Seg)
		var neighbours = tree.Search(s.BBox())

		for _, node := range neighbours {
			var o = node.GetItem().(*seg.Seg)
			if s == o {
				continue
			}

			var k = cacheKey(s, o)
			if cache[k] {
				continue
			}
			cache[k] = true

			var intersects = s.Segment.SegSegIntersection(o.Segment)
			for _, pt := range intersects {
				if pt.IsVertex() && !pt.IsVerteXOR() { //if not exclusive vertex
					continue
				}
				cg := ctx.New(pt.Point, 0, -1).AsNonPlanarVertex()
				cg.Meta.NonPlanar = iter.SortedIntsSet(k[:])
				results.Push(cg)
			}
		}
	}
	return results
}

//cache key: [0, 1, 9, 10] == [9, 10, 0, 1]
func cacheKey(a, b *seg.Seg) [4]int {
	if b.I < a.I {
		a, b = b, a
	}
	return [4]int{a.I, a.J, b.I, b.J}
}

func segmentDB(polyline *pln.Polyline) (*rtree.RTree, []rtree.BoxObj) {
	var tree = rtree.NewRTree(4)
	var data = make([]rtree.BoxObj, 0)
	for _, s := range polyline.Segments() {
		data = append(data, s)
	}
	tree.Load(data)
	return tree, data
}
