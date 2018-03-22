package lnr

import (
	"simplex/pln"
	"github.com/intdxdt/rtree"
)

func SegmentDB(polyline *pln.Polyline, nodeCapacity ...int ) (*rtree.RTree, []rtree.BoxObj) {
	var capacity = 4
	if len(nodeCapacity) > 0 {
		capacity = nodeCapacity[0]
	}

	var tree = rtree.NewRTree(capacity)
	var data = make([]rtree.BoxObj, 0)
	for _, s := range polyline.Segments() {
		data = append(data, s)
	}
	tree.Load(data)

	return tree, data
}
