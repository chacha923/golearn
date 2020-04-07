package list

//翻转链表
func Reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next
	pre := head
	pre.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

//判断链表是否有环
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	var slow *ListNode //快慢指针
	var fast *ListNode
	fast = head.Next
	slow = head
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

//得到链表倒数第n个节点
func nthToLast(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return nil
	}
	l1 := head //两个指针, 一个先前进n-1步
	l2 := head

	for i := 0; i < n-1; i++ {
		if l2 == nil {
			return nil
		}
		l2 = l2.Next
	}

	for l2.Next != nil {
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1
}
