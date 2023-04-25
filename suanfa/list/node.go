package list

// 基础的链表节点
type Node struct {
	// 数据域
	Val int
	// 指针域
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

type DoubleNode struct {
	// 数据域
	Val int
	// 指针域
	Next *DoubleNode // 后一个
	Prev *DoubleNode // 前一个
}

func NewDoubleNode(val int) *DoubleNode {
	return &DoubleNode{Val: val}
}
