package lnr

import (
	"sort"
	"github.com/intdxdt/geom"
)


type vertex struct {
	*geom.Point
	index int
}

type vertices []*vertex

//lexical sort of x and y coordinates
func (v vertices) Less(i, j int) bool {
	return (v[i].Point[x] < v[j].Point[x]) || (
			v[i].Point[x] == v[j].Point[x] &&
			v[i].Point[y] < v[j].Point[y])
}

//Len for sort interface
func (v vertices) Len() int {
	return len(v)
}

//Swap for sort interface
func (v vertices) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

//Inplace Lexicographic sort
func (v vertices) Sort() {
	sort.Sort(v)
}
