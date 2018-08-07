package lnr

import (
	"github.com/intdxdt/iter"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/rtree"
	"github.com/TopoSimplify/ctx"
	"github.com/TopoSimplify/pln"
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
	var points = make([]vertex, 0, polyline.Coordinates.Len())
	for i := range polyline.Coordinates.Idxs {
		points = append(points, vertex{polyline.Pt(i), i, NullFId})
	}
	vertices(points).Sort() //O(nlogn)

	var d = 0
	var a, b *vertex
	var indices []int
	var results = ctx.NewContexts()

	var bln bool
	for i, n := 0, len(points); i < n-1; i++ { //O(n)
		a, b = &points[i], &points[i+1]
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
			var cg = ctx.New(points[i].Point, 0, -1).AsPlanarVertex()
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
	var s *geom.Segment
	var neighbours []*rtree.Obj

	for _, d := range data {
		s = d.Object.(*geom.Segment)
		neighbours = tree.Search(s.Bounds())

		for _, obj := range neighbours {
			var o = obj.Object.(*geom.Segment)
			if s == o {
				continue
			}

			var k = cacheKey(s, o)
			if cache[k] {
				continue
			}
			cache[k] = true

			var intersects = s.SegSegIntersection(o)
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
func cacheKey(a, b *geom.Segment) [4]int {
	if b.Coords.Idxs[0] < a.Coords.Idxs[0] {
		a, b = b, a
	}
	return [4]int{a.Coords.Idxs[0], a.Coords.Idxs[1], b.Coords.Idxs[0], b.Coords.Idxs[1]}
}

func segmentDB(polyline *pln.Polyline) (*rtree.RTree, []*rtree.Obj) {
	var tree = rtree.NewRTree(4)
	var data = make([]*rtree.Obj, 0)
	var segments = polyline.Segments()
	for i := range segments {
		data = append(data, rtree.Object(i, segments[i].Bounds(), segments[i]))
	}
	tree.Load(data)
	return tree, data
}
