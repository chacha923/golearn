package list

import (
	"fmt"
	"math"
)

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

func NewEmptyNode() *Node {
	return &Node{Val: math.MinInt64}
}

// 追加节点, 用于构建链表
func (cls *Node) Append(val int) {
	if cls == nil {
		return
	}
	for cls.Next != nil {
		cls = cls.Next
	}
	cls.Next = NewNode(val)
}

// 打印
func (cls *Node) Print() {
	if cls == nil {
		return
	}
	if cls.Next == nil {
		fmt.Println(cls.Val)
	}
	for cls != nil {
		fmt.Printf("%d", cls.Val)
		if cls.Next != nil {
			fmt.Printf(" -> ")
		}
		cls = cls.Next
	}
	fmt.Println()
}

// 当前节点和 next 节点，反转指针域
// 1 -> 2 => 2 -> 1
// 这是一个简单实现，会导致链表断裂，需要实现保存当前节点，不能直接用
func (cls *Node) Reverse(val int) {
	if cls == nil {
		return
	}
	if cls.Next == nil {
		return
	}
	var temp *Node
	temp = cls
	cls = cls.Next
	cls.Next = temp
}
