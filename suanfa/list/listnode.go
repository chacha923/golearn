package list

type ListNode struct {
	Next *ListNode
	Value int
}

func NewListNode(value int) *ListNode{
	return &ListNode{Value:value}
}
