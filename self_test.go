package lnr

import (
	"time"
	"testing"
	"simplex/pln"
	"github.com/franela/goblin"
	"github.com/intdxdt/geom"
)

func newPolyline(wkt string) *pln.Polyline {
	return pln.New(geom.NewLineStringFromWKT(wkt).Coordinates())
}

func TestToSelfIntersects(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("planar self intersection", func() {
		g.It("should test constrain to self intersects", func() {
			g.Timeout(1 * time.Hour)
			var ln = newPolyline("LINESTRING ( 740 380, 720 440, 760 460, 740 520, 860 520, 860 620, 740 620, 740 520, 640 520, 640 420, 841 420, 840 320 )")
			var inters = SelfIntersection(ln)

			g.Assert(inters.Len()).Equal(2)
			ln = newPolyline("LINESTRING ( 1000 600, 1100 600, 1100 500, 1000 500, 1000 400, 1100 400, 1100 500, 1200 500, 1200 400, 1300 400, 1300 500, 1200 500, 1200 600, 1100 600 )")
			inters = SelfIntersection(ln)
			g.Assert(inters.Len()).Equal(3)

			ln = newPolyline("LINESTRING ( 1100 100, 1300 300, 1400 200, 1400 100, 900 100, 900 0, 1100 0, 1100 100, 1000 300, 900 200, 1100 100, 1300 0, 1200 -100, 1100 -100, 1100 -200, 1300 -200, 1300 0 )")
			inters = SelfIntersection(ln)
			g.Assert(inters.Len()).Equal(2)
		})
	})
}
