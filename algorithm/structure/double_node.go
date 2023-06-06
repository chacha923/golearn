package structure

import "math"

// 双向链表节点
type DoubleNode struct {
	// 数据域
	Key string // 键值对
	Val int
	// 指针域
	Next *DoubleNode // 后一个
	Prev *DoubleNode // 前一个
}

func NewDoubleNode(val int) *DoubleNode {
	return &DoubleNode{Val: val}
}

func NewDoubleNodeWithKey(key string, val int) *DoubleNode {
	return &DoubleNode{Key: key, Val: val}
}

func NewEmptyDoubleNode() *DoubleNode {
	return &DoubleNode{Val: math.MinInt64}
}

// 带使用频率的双端队列节点, lfu
type FreqDequeNode struct {
	Freq int
	DoubleNode
}
