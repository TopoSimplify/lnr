package lnr

import (
	"github.com/intdxdt/sset"
	"github.com/intdxdt/cmp"
)

func FeatureClassSelfIntersection(featureClass []Linear) map[string]*sset.SSet {
	var dict = make(map[[2]float64]map[string]int, 0)
	for _, self := range featureClass {
		polyline := self.Polyline()
		n := polyline.Len()
		for i := 0; i < n; i++ {
			var dat map[string]int
			var ok bool
			var pt = polyline.Coordinates[i]
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
