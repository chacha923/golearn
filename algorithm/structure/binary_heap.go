package structure

import (
	"fmt"
	"golearn/algorithm/util"
)

// 二叉堆

// 二叉堆结构体
// 二叉堆节点结构体
type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

type BinaryHeap struct {
	root *Node
}

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{
		root: nil,
	}
}

// 向上调整堆，保持堆的性质
func (h *BinaryHeap) heapifyUp(node *Node) {
	if node == nil || node == h.root {
		return
	}
	var parent = node.parent
	// 小顶堆，父亲应该比儿子小
	if node.value < parent.value {
		util.Swap(&node.value, &parent.value)
		// 父亲变小了，但还要向上查看直到堆顶
		h.heapifyUp(parent)
	}
}

// 向下调整堆，保持堆的性质
func (h *BinaryHeap) heapifyDown(node *Node) {
	if node == nil {
		return
	}
	// 假设父亲是最小的，如果不是，那么就要和最小的儿子交换，再向下查看
	var min = node
	if node.left != nil && node.left.value < min.value {
		min = node.left
	}
	if node.right != nil && node.right.value < min.value {
		min = node.right
	}
	if min != node {
		util.Swap(&min.value, &node.value)
		h.heapifyDown(min)
	}
}

// 插入一个元素到堆中
func (h *BinaryHeap) Insert(value int) {

}

// 删除堆顶元素
func (h *BinaryHeap) DeleteMin() {
}

// 删除指定节点
func (h *BinaryHeap) deleteNode(node *Node) {
}

// 获取父节点
func (h *BinaryHeap) getParent(node *Node) *Node {
	if node == nil || node == h.root {
		return nil
	}

	return node.parent
}

// 获取堆顶元素
func (h *BinaryHeap) GetMin() int {
	if h.root == nil {
		return -1 // 表示堆为空
	}
	return h.root.value
}

// 打印堆中的元素
func (h *BinaryHeap) PrintHeap() {
	if h.root == nil {
		fmt.Println("Heap is empty")
		return
	}

	queue := []*Node{h.root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", curr.value)

		if curr.left != nil {
			queue = append(queue, curr.left)
		}

		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}

	fmt.Println()
}
