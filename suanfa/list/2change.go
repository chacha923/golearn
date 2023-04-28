package list

// 两两交换节点
func swapPairs(head *Node) *Node {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var dummy = NewEmptyNode()
	var temp = dummy

	for {
		// 终止条件，只剩一个节点或者没有节点
		if temp.Next == nil || temp.Next.Next == nil {
			break
		}
		var node1 = temp.Next
		var node2 = temp.Next.Next

		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		temp = node1

	}
	return dummy.Next
}

// 反转链表节点，只需要知道前序节点，链表就不会断
// temp		temp.next  temp.next.next
// temp	->	node1  -> node2
