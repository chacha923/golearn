package structure

import (
	"crypto/sha1"
	"math"
	"sort"
	"strconv"
	"sync"
)

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
	virtualSpots int
	nodes        nodeArray
	weights      map[string]int
	mu           sync.RWMutex
}

func NewHashRing(spots int) *HashRing {
	if spots == 0 {
		spots = DefaultVirtualSports
	}
	h := &HashRing{
		virtualSpots: spots,
		weights:      make(map[string]int),
	}
	return h
}

func (h *HashRing) AddNodes(nodes map[string]int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for key, weight := range nodes {
		h.weights[key] = weight
	}
	h.generate()
}

func (h *HashRing) RemoveNode(key string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.weights, key)
	h.generate()
}

func (h *HashRing) UpdateNode(key string, weight int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.weights[key] = weight
}

//GetNode get node with key
func (h *HashRing) GetNode(s string) string {
	h.mu.Lock()
	defer h.mu.Unlock()
	if len(h.nodes) == 0 {
		return ""
	}

	hash := sha1.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)
	v := genValue(hashBytes[6:10])
	i := sort.Search(len(h.nodes), func(i int) bool {
		return h.nodes[i].spotValue >= v
	})
	if i == len(h.nodes) {
		i = 0
	}
	return h.nodes[i].nodeKey
}

func (h *HashRing) generate() {
	var totalW int
	for _, w := range h.weights {
		totalW += w
	}

	totalVirtualSpots := h.virtualSpots * len(h.weights)
	h.nodes = nodeArray{}

	for key, weight := range h.weights {
		spots := int(math.Floor(float64(weight) / float64(totalW) * float64(totalVirtualSpots)))
		for i := 1; i <= spots; i++ {
			hash := sha1.New()
			hash.Write([]byte("node key" + ":" + strconv.Itoa(i)))
			hashBytes := hash.Sum(nil)
			n := node{
				nodeKey:   key,
				spotValue: genValue(hashBytes[6:10]),
			}
			h.nodes = append(h.nodes, n)
			hash.Reset()
		}
	}
	h.nodes.Sort()
}

func genValue(bs []byte) uint32 {
	if len(bs) < 4 {
		return 0
	}
	v := (uint32(bs[3]) << 24) | (uint32(bs[2]) << 16) | (uint32(bs[1]<<8) | (uint32(bs[0])))
	return v
}
