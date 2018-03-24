package lnr

import "github.com/intdxdt/geom"

type interPt struct {
	*geom.Point
	inter int
}

func (p *interPt) isCollinear() bool {
	return p.inter == collinear
}

func (p *interPt) isSegmentIntersection() bool {
	return p.inter == seginters
}

func (p *interPt) isVertexIntersection() bool {
	return p.inter == vertexinters
}