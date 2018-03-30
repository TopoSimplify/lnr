package main

import "github.com/intdxdt/geom"

type interPt struct {
	*geom.Point
	inter int
}

func (p *interPt) isCollinear() bool {
	return p.inter == collinear
}

func (p *interPt) isSegmentIntersection() bool {
	return p.inter == segmentIntersect
}

func (p *interPt) isVertexIntersection() bool {
	return p.inter == vertexIntersect
}