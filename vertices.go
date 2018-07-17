package lnr

import (
	"sort"
	"github.com/intdxdt/geom"
)

type vertex struct {
	geom.Point
	index int
	fid   string
}

type vertices []*vertex

//lexical sort of x and y coordinates
func (v vertices) Less(i, j int) bool {
	return (v[i].Point[x] <  v[j].Point[x]) || (
			v[i].Point[x] == v[j].Point[x] &&
			v[i].Point[y] <  v[j].Point[y])
}

//Len for sort interface
func (v vertices) Len() int {
	return len(v)
}

//Swap for sort interface
func (v vertices) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

//In place lexicographic sort
func (v vertices) Sort() {
	sort.Sort(v)
}


func appendVertices(points []*vertex, coordinates []geom.Point, fid string) []*vertex {
	for i := range coordinates {
		points = append(points, &vertex{coordinates[i], i, fid})
	}
	return points
}
