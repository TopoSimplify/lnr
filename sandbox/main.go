package main

import (
	"fmt"
	"github.com/intdxdt/geom"
)

func main() {
	var wkt = "LINESTRING ( 2476.199991591663 -185.7000126125052, 2681.703400801933 124.33277809745438 )"
	var as = geom.NewSegment(geom.NewLineStringFromWKT(wkt).Coordinates, 0, 1)

	var bwkt = "LINESTRING ( 2400 -300, 2600 0, 2900 100 )"
	var bs = geom.NewSegment(geom.NewLineStringFromWKT(bwkt).Coordinates, 0, 1)

	var inters = as.SegSegIntersection(bs)
	for _, p := range inters {
		fmt.Println(p.WKT())
	}
}
