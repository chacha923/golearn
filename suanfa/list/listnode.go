package list

type ListNode struct {
	Next *ListNode
	Val  int
}

type DequeNode struct {
	Key  int
	Val  int
	Pre  *DequeNode
	Next *DequeNode
}

func NewListNode(Val int) *ListNode {
	return &ListNode{Val: Val}
}
