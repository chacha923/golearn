package hash

import "sort"

const (
	DefaultVirtualSports = 400
)

type node struct {
	nodeKey   string
	spotValue uint32
}

type nodeArray []node

func (p nodeArray) Len() int {
	return len(p)
}

func (p nodeArray) Less(i, j int) bool {
	return p[i].spotValue < p[j].spotValue
}

func (p nodeArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p nodeArray) Sort() {
	sort.Sort(p)
}

type HashRing struct {
}
