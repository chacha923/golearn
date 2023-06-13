package list

// 链表的 partition 函数，把原链表一分为二，一个链表中的值都小于 x，另一个链表中的值都大于等于 x
// 最后把这两个链表连接起来返回新的链表的头节点
func partition(head *Node, x int) *Node {
	// 小于 x
	var dummy1 = NewEmptyNode()
	// 大于等于 x
	var dummy2 = NewEmptyNode()

	var p1 = dummy1
	var p2 = dummy2

	var p = head // p 负责遍历原链表

	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else if p.Val >= x {
			p2.Next = p
			p2 = p2.Next
		}
		// 往前移动，并断开当前节点 next 指针域，
		var temp = p.Next
		p.Next = nil
		p = temp
	}
	// 将两个链表连接起来
	p1.Next = dummy2.Next

	return dummy1.Next
}
