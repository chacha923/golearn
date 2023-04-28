package list

// K个一组翻转链表
// 给你这个链表：1->2->3->4->5
// 当 k = 2 时，应当返回: 2->1->4->3->5
// 当 k = 3 时，应当返回: 3->2->1->4->5
// 原地反转, 空间复杂度O(1)
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/tu-jie-kge-yi-zu-fan-zhuan-lian-biao-by-user7208t/
func reverseKGroup(head *Node, k int) *Node {
	if head == nil {
		return nil
	}
	dummy := NewEmptyNode()
	dummy.Next = head // dummy next 永远指向头结点
	// 需要两个临时指针，跨度为 k
	pre := dummy
	end := dummy

	for end.Next != nil {
		// 找出本组最后一个节点
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		// end 走到底, 可以返回了
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next          // 保存当前链表组(k个一组)的下一个节点
		end.Next = nil            // 截断k个节点
		pre.Next = reverse(start) // 在当前组执行反转
		start.Next = next         // 反转后, start指向的节点跑到最后, 指向next

		// 指向当前组最后一个节点, 即下一组头的前一个节点
		pre = start
		end = pre
	}
	// 为啥dummy一定指向头节点? 第一次执行pre和dummy指向同一个节点, pre.Next修改, dummy.Next也就修改了
	return dummy.Next
}

// 用栈实现k个一组翻转链表, 空间复杂度O(k), 用改变链表指针域的方式
func reverseKGroupWithStack(head *Node, k int) *Node {
	if head == nil {
		return nil
	}
	dummyHead := &Node{Next: head} // next域永远指向当前头节点
	preNode := dummyHead           // 临时指针
	stack := make([]*Node, 0)      // 临时栈

	for head != nil {
		count := 0
		// head 移动k次
		for head != nil && count < k {
			stack = append(stack, head)
			head = head.Next
			count += 1
		}
		// 长度不够, 不反转
		if k > len(stack) {
			preNode.Next = stack[0]
			break
		}
		for i := len(stack) - 1; i >= 0; i-- {
			preNode.Next = stack[i]
			preNode = stack[i] // 前进
		}
		preNode.Next = head // 反转好的部分接上原链表
		stack = stack[0:0]  // 清空栈
	}
	return dummyHead.Next
}

// k个一组翻转链表, 直接修改节点值
// 比较作弊，一般不允许改变值域
func reverseKGroupWithChangeValue(head *Node, k int) *Node {
	if head == nil {
		return nil
	}
	stack := make([]int, 0) // 临时栈
	dummy := NewEmptyNode()
	dummy.Next = head
	cur := head
	for cur != nil {
		count := 0
		for cur != nil && count < k {
			stack = append(stack, cur.Val)
			cur = cur.Next
			count += 1
		}
		// 特殊情况
		if k > len(stack) {
			break
		}
		for i := len(stack) - 1; i >= 0; i-- {
			head.Val = stack[i]
			head = head.Next
		}
		stack = stack[0:0]
	}
	return dummy.Next
}
