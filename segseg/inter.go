package lnr

import (
	"github.com/intdxdt/geom"
)

type IntPt struct {
	*geom.Point
	Inter VBits
}

type IntPts []*IntPt

func (s IntPts) Len() int {
	return len(s)
}
func (s IntPts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IntPts) Less(i, j int) bool {
	return lexsort2d(s[i], s[j]) < 0
}

func (p *IntPt) IsIntersection() bool {
	return p.Inter == 0
}

func (p *IntPt) IsVertex() bool {
	var mask = SelfMask |OtherMask
	return p.Inter& mask  > 0
}
