package list

type ListNode struct {
	Next  *ListNode
	Value int
}

type DequeNode struct {
	Key   int
	Value int
	Pre   *DequeNode
	Next  *DequeNode
}

func NewListNode(value int) *ListNode {
	return &ListNode{Value: value}
}
