package lnr

import (
	"github.com/intdxdt/geom"
	"random"
)

type FC struct {
	Coordinates []*geom.Point
	Fid         string
}

func NewFC(coordinates []*geom.Point) *FC {
	return &FC{coordinates, random.String(10)}
}

func (fc *FC) Id() string {
	return fc.Fid
}
