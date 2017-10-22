package lnr

import (
	"simplex/ctx"
	"simplex/pln"
	"simplex/seg"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/sset"
	"github.com/intdxdt/rtree"
)

type kvCount struct {
	count   int
	indxset *sset.SSet
}

type selfInter struct {
	keyset *sset.SSet
	point  *geom.Point
}

func updateKVCount(dict map[[2]float64]*kvCount, o *geom.Point, index int) {
	var k = [2]float64{o[0], o[1]}
	var v, ok = dict[k]
	if !ok {
		v = &kvCount{
			count:   0,
			indxset: sset.NewSSet(cmp.Int, 8),
		}
		dict[k] = v
	}
	v.indxset.Add(index)
	v.count += 1
}

func SelfIntersection(polyline *pln.Polyline) *ctx.ContextGeometries {
	var tree = *rtree.NewRTree(16)
	var dict = make(map[[2]float64]*kvCount)
	var data = make([]rtree.BoxObj, 0)

	for _, s := range polyline.Segments() {
		data = append(data, s)
	}
	tree.Load(data)

	var selfIntersects = make(map[string]*selfInter, 0)

	for _, d := range data {
		var s = d.(*seg.Seg)
		var res = tree.Search(s.BBox())
		updateKVCount(dict, s.A, s.I)
		updateKVCount(dict, s.B, s.J)

		for _, node := range res {
			var otherSeg = node.GetItem().(*seg.Seg)
			if s == otherSeg {
				continue
			}

			var segG, otherSegG = s.Segment, otherSeg.Segment
			var intersects = segG.Intersection(otherSegG)

			if len(intersects) == 0 {
				continue
			}

			for _, pt := range intersects {
				if s.A.Equals2D(pt) || s.B.Equals2D(pt) {
					continue
				}
				var skey = sset.NewSSet(cmp.Int, 8).Extend(
					s.I, s.J, otherSeg.I, otherSeg.J,
				)

				var k = skey.String()
				v, ok := selfIntersects[k]
				if !ok {
					v = &selfInter{keyset: skey, point: pt}
				}
				selfIntersects[k] = v
			}
		}
	}

	var results = ctx.NewContexts()
	for _, val := range selfIntersects {
		cg := ctx.New(val.point, 0, -1).AsSelfNonVertex()
		cg.Meta.SelfNonVertices = val.keyset
		results.Push(cg)
	}

	for k, v := range dict {
		if v.count > 2 {
			i := v.indxset.First().(int)
			cg := ctx.New(geom.NewPoint(k[:]), i, i).AsSelfVertex()
			cg.Meta.SelfVertices = v.indxset
			results.Push(cg)
		}
	}

	return results.Sort()
}
