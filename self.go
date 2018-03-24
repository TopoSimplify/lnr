package lnr

import (
	"simplex/ctx"
	"simplex/seg"
	"simplex/pln"
	"github.com/intdxdt/iter"
	"github.com/intdxdt/sset"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/geom"
)
const emptyFId = ""

//update planar and non-planar intersections
func SelfIntersection(polyline *pln.Polyline) *ctx.ContextGeometries {
	var results = ctx.NewContexts()
	var planar = planarIntersects(polyline).DataView()
	var non_planar = nonPlanarIntersection(polyline).DataView()
	return results.Extend(planar).Extend(non_planar)
}

//Planar self-intersection
func planarIntersects(polyline *pln.Polyline) *ctx.ContextGeometries {
	var points = make([]*vertex, 0, len(polyline.Coordinates))
	for i, pt := range polyline.Coordinates {
		points = append(points, &vertex{pt, i,emptyFId})
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
	var tree, data = SegmentDB(polyline)
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

func FeatureClassSelfIntersection(featureClass []*FC) map[string]*sset.SSet {
	var dict = make(map[[2]float64]map[string]int, 0)
	for _, self := range featureClass {
		var n = len(self.Coordinates)

		for i := 0; i < n; i++ {
			var ok bool
			var dat map[string]int
			var pt = self.Coordinates[i]
			var key = [2]float64{pt.X(), pt.Y()}

			if dat, ok = dict[key]; !ok {
				dat = make(map[string]int, 0)
			}
			var id = self.Id()
			if _, ok = dat[id]; !ok {
				dat[id] = i
			}
			dict[key] = dat
		}
	}

	var junctions = make(map[string]*sset.SSet, 0)
	for _, o := range dict {
		if len(o) > 1 {
			for sid, idx := range o {
				var ok bool
				var s *sset.SSet

				if s, ok = junctions [sid]; !ok {
					s = sset.NewSSet(cmp.Int)
				}
				s.Add(idx)
				junctions[sid] = s
			}
		}
	}
	return junctions
}

//Planar self-intersection
func FeatureClassPlanarIntersects(featureClass []*FC) map[string][]int{
	//preallocate points size
	var n = 0
	for _, self := range featureClass {
		n += len(self.Coordinates)
	}
	var points = make([]*vertex, 0, n)

	for _, self := range featureClass {
		points = appendPoints(points, self.Coordinates, self.Id())
	}

	vertices(points).Sort() //O(nlogn)

	var d = 0
	var indexes []*vertex
	var results = make(map[string][]int, 0)
	var a, b *vertex

	var bln bool
	for i, n := 0, len(points); i < n-1; i++ { //O(n)
		a, b = points[i], points[i+1]
		bln = a.Equals2D(b.Point)
		if bln {
			if a.fid == b.fid {
				continue
			}
			if d == 0 {
				indexes = append(indexes, a, b)
				d += 2
			} else {
				indexes = append(indexes, b)
				d += 1
			}
			continue
		}

		if d > 1 {
			for _, v := range indexes {
				results[v.fid] = append(results[v.fid], v.index)
			}
		}
		d = 0
		indexes = indexes[:0]
	}
	return results
}

func appendPoints(points []*vertex, coordinates []*geom.Point, fid string) []*vertex {
	for i, pt := range coordinates {
		points = append(points, &vertex{pt, i, fid})
	}
	return points
}
