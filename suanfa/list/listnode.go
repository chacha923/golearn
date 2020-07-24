package list

// 各种链表节点定义

// 链表节点
type ListNode struct {
	Next *ListNode
	Val  int
}

// 双端队列节点
type DequeNode struct {
	Key  int
	Val  int
	Pre  *DequeNode
	Next *DequeNode
}

// 带使用频率的双端队列节点, lfu
type FreqDequeNode struct {
	Freq int
	DequeNode
}

func NewListNode(Val int) *ListNode {
	return &ListNode{Val: Val}
}
