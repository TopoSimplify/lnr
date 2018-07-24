package main

import (
	"fmt"
	"github.com/intdxdt/geom"
	"github.com/TopoSimplify/lnr"
)

func main() {
	var wkt = "LINESTRING ( 2476.199991591663 -185.7000126125052, 2681.703400801933 124.33277809745438 )"
	var ac = geom.NewLineStringFromWKT(wkt).Coordinates()
	var as = geom.NewSegment(&ac[0], &ac[1])

	var bwkt = "LINESTRING ( 2400 -300, 2600 0, 2900 100 )"
	var bc = geom.NewLineStringFromWKT(bwkt).Coordinates()
	var bs = geom.NewSegment(&bc[0], &bc[1])

	var inters = lnr.SegIntersection(as,bs)
	for _, p := range inters {
		fmt.Println(p.WKT())
	}
}
