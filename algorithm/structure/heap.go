package structure

// 实现一个普通的堆，用到  container/heap interface
type Heap []int

// 实现 heap.Interface 接口的方法

// Len is the number of elements in the collection.
func (h Heap) Len() int {
	return len(h)
}

// Less reports whether the element with
// index i should sort before the element with index j.
// 小顶堆
func (h Heap) Less(i, j int) bool {
	// 此处为小根堆，如果要实现大根堆，修改为 h[i] > h[j]
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x int) {
	*h = append(*h, x)
}

func (h *Heap) Pop() int {
	var old = *h
	var n = old.Len()
	var x = old[n-1]
	*h = old[0 : n-1]
	return x
}
