package list

// 返回倒数第k个节点
func kthToLast(head *ListNode, k int) int {
	var dummy = &ListNode{}
	dummy.Next = head
	var fast, slow = dummy, dummy

	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}
	for {
		if fast.Next == nil {
			return slow.Value
		}
		fast = fast.Next
		slow = slow.Next
	}
}
