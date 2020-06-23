package list

import (
	"testing"
	"fmt"
)

var head *ListNode
func init() {
	head = NewListNode(1)
	head.Next = NewListNode(2)
	head.Next.Next = NewListNode(3)
	head.Next.Next.Next = NewListNode(4)
	head.Next.Next.Next.Next = NewListNode(5)
	head.Next.Next.Next.Next.Next = NewListNode(6)
	head.Next.Next.Next.Next.Next.Next = NewListNode(7)
}

func TestReverse(t *testing.T) {
	newHead := Reverse(head)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}
