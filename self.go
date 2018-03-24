package lnr

import (
	"simplex/ctx"
	"simplex/seg"
	"simplex/pln"
	"github.com/intdxdt/iter"
	"github.com/intdxdt/rtree"
)

//update planar and non-planar intersections
func SelfIntersection(polyline *pln.Polyline) *ctx.ContextGeometries {
	return planarIntersects(polyline).Extend(
		nonPlanarIntersection(polyline).DataView(),
	)
}

//Planar self-intersection
func planarIntersects(polyline *pln.Polyline) *ctx.ContextGeometries {
	var points = make([]*vertex, 0, len(polyline.Coordinates))
	for i, pt := range polyline.Coordinates {
		points = append(points, &vertex{pt, i, NullFId})
	}
	vertices(points).Sort() //O(nlogn)
	var results = ctx.NewContexts()
	var d = 0
	var indices []int
	var a, b *vertex

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
			cg.Meta.Planar = append([]int{}, indices...)
			iter.SortedIntSet(&cg.Meta.Planar)
			results.Push(cg)
		}
		d = 0
		indices = indices[:0]
	}
	return results
}

func nonPlanarIntersection(polyline *pln.Polyline) *ctx.ContextGeometries {
	var dict = make(map[[4]int]struct{})
	var tree, data = segmentDB(polyline)
	var results = ctx.NewContexts()
	//update planar and non-planar dictionaries
	for _, d := range data {
		var s = d.(*seg.Seg)
		var neighbours = tree.Search(s.BBox())

		for _, node := range neighbours {
			var o = node.GetItem().(*seg.Seg)
			if s == o {
				continue
			}

			var intersects = SegIntersection(s.Segment, o.Segment)

			for _, pt := range intersects {
				if pt.isVertexIntersection() {
					continue
				}

				var k = [4]int{s.I, s.J, o.I, o.J}
				var key = k[:]
				iter.SortedIntSet(&key)

				if _, ok := dict[k]; !ok {
					var cg = ctx.New(pt.Point, 0, -1).AsNonPlanarVertex()
					cg.Meta.NonPlanar = key
					results.Push(cg)
				}
				dict[k] = struct{}{}
			}
		}
	}
	return results
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

