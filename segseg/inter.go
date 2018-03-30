package lnr

import (
	"strings"
	"github.com/intdxdt/geom"
)

type IntPt struct {
	*geom.Point
	inter string
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

func (p *IntPt) IsCollinear() bool {
	return p.inter[0] == 'c'
}

func (p *IntPt) IsIntersection() bool {
	return p.inter[0] == 'x'
}

func (p *IntPt) IsVertex() bool {
	return p.inter[0] == 'v'
}

func (p *IntPt) Tokens() []string {
	return strings.Split(p.inter, "-")
}
