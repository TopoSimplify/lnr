package lnr

import (
	"simplex/ctx"
	"simplex/seg"
	"simplex/pln"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/sset"
	"github.com/intdxdt/rtree"
)

type vertexDegree struct {
	degree   int
	indexSet *sset.SSet
}

type crossings struct {
	keySet *sset.SSet
	point  *geom.Point
}

func SelfIntersection(polyline *pln.Polyline) *ctx.ContextGeometries {

	var nonPlanarDict = make(map[string]*crossings, 0)
	var planarDict = make(map[[2]float64]*vertexDegree)
	var tree, data = SegmentDB(polyline)

	//update planar and non-planar dictionaries
	for _, d := range data {
		var s = d.(*seg.Seg)
		updatePlanarDict(planarDict, s)
		updateNonPlanarDict(nonPlanarDict, s, tree)
	}

	var results = ctx.NewContexts()

	for _, val := range nonPlanarDict {
		var cg = ctx.New(val.point, 0, -1).AsNonPlanarVertex()
		cg.Meta.NonPlanarVertices = val.keySet
		results.Push(cg)
	}

	for k, v := range planarDict {
		if v.degree > 2 {
			var cg = ctx.New(geom.NewPoint(k[:]), 0, -1).AsPlanarVertex()
			cg.Meta.PlanarVertices = v.indexSet
			results.Push(cg)
		}
	}

	return results
}

func updatePlanarDict(dict map[[2]float64]*vertexDegree, s *seg.Seg) {
	updateVertexDegree(dict, s.A, s.I)
	updateVertexDegree(dict, s.B, s.J)
}

func updateNonPlanarDict(dict map[string]*crossings, s *seg.Seg, tree *rtree.RTree) {
	var neighbours = tree.Search(s.BBox())
	for _, node := range neighbours {
		var o = node.GetItem().(*seg.Seg)
		if s == o {
			continue
		}

		var intersects = s.Intersection(o.Segment)

		if len(intersects) == 0 {
			continue
		}

		for _, pt := range intersects {
			if s.A.Equals2D(pt) || s.B.Equals2D(pt) {
				continue
			}
			var key = sset.NewSSet(cmp.Int, 8).Extend(s.I, s.J, o.I, o.J)
			var k = key.String()
			v, ok := dict[k]
			if !ok {
				v = &crossings{keySet: key, point: pt}
			}
			dict[k] = v
		}
	}
}

func updateVertexDegree(dict map[[2]float64]*vertexDegree, o *geom.Point, index int) {
	var k = [2]float64{o[0], o[1]}
	var v, ok = dict[k]
	if !ok {
		v = &vertexDegree{
			degree:   0,
			indexSet: sset.NewSSet(cmp.Int, 8),
		}
		dict[k] = v
	}
	v.indexSet.Add(index)
	v.degree += 1
}
