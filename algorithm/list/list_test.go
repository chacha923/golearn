package list

import (
	"fmt"
	"testing"
)

var head *Node

func init() {
	head = NewNode(1)
	head.Append(2)
	head.Append(3)
	head.Append(4)
	head.Append(5)
	head.Append(6)
	head.Append(7)
}

func TestReverse(t *testing.T) {
	newHead := Reverse(head)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}

func TestDeleteDuplicates2(t *testing.T) {
	// 构造链表
	head := NewNode(1)
	head.Append(1)
	head.Append(1)
	head.Append(2)
	head.Append(5)

	head.Print()

	var newHead = deleteDuplicates2(head)
	fmt.Println("result1")
	newHead.Print()

	head = NewNode(1)
	head.Append(2)
	head.Append(3)
	head.Append(3)
	head.Append(4)
	head.Append(4)
	head.Append(5)
	head.Print()

	newHead = deleteDuplicates2(head)
	fmt.Println("result2")
	newHead.Print()

}
