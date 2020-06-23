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
		next := cur.Next // 保存下一个节点
		cur.Next = pre   // 反转next域
		pre = cur        // 往前走
		cur = next
	}
	return pre
}

// 判断链表是否有环
// 快指针追上慢指针说明有环
// 快指针走到链表尾说明无环
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

// 返回倒数第k个节点,  使用dummy节点
func kthToLast(head *ListNode, k int) int {
	if k <= 0 || head == nil {
		return -1
	}
	var dummy = &ListNode{}
	dummy.Next = head             // dummy指向头节点
	var fast, slow = dummy, dummy // 快慢指针

	// 快指针先走k-1步
	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}
	for {
		if fast.Next == nil {
			return slow.Val
		}
		fast = fast.Next
		slow = slow.Next
	}
}

// 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
// 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
// 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
// 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
// 首先用快慢指针找到奇数部分最后一个节点（总和奇数偶数情况），再从头开始，一次遍历到链表尾部，将尾部插到当前节点后面
// 可以不需要第一部利用快慢指针找最后一个节点，这里这么做的目的，当节点数特别大的时候，可以节约一半的遍历时间
// https://leetcode-cn.com/problems/reorder-list/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-34/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// 快慢指针切割链表为前后两部分
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	left := head       // 前半部分头结点
	right := slow.Next // 后半部分头结点
	slow.Next = nil
	right = reverse(right) // 反转后半部分

	// 合并
	for right != nil {
		lNext := left.Next  // 存left next节点
		rNext := right.Next // 存 right next节点
		left.Next = right   // left next 指向right
		right.Next = lNext  // right next指向left第二个节点
		left = lNext        // 移动头指针left right
		right = rNext
	}
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// K个一组翻转链表
// 给你这个链表：1->2->3->4->5
// 当 k = 2 时，应当返回: 2->1->4->3->5
// 当 k = 3 时，应当返回: 3->2->1->4->5
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	demmy := &ListNode{}
	demmy.Next = head // dummy next 永远指向头结点
	pre := dummy
	curr := dummy
	end := dummy
	for end != nil {
		// 当前
		curr := pre.Next
		// 结束节点
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil {
				return dummy.Next
			}
		}
	}

}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

// 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
func sortList(head *ListNode) *ListNode {
	// 如果 head为空或者head就一位,直接返回
	if head == nil || head.Next == nil {
		return head
	}
	// 定义快慢俩指针,当快指针到末尾的时候,慢指针肯定在链表中间位置
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 把链表拆分成两段,所以设置中间位置即慢指针的next为nil
	n := slow.Next  // 后一半链表头
	slow.Next = nil // 中点截断, 产生两个链表
	// 递归排序
	return MergeLists2(sortList(head), sortList(n))
}

// 找出两个链表相交的起始节点 (公共节点)
func GetIntersectionNode(headA, headB *ListNode) {
	if headA == nil || headB == nil {
		return nil
	}
	// 相交的长度相等
	lenA, lenB := 0, 0
	// 第一次遍历找长度
	for p := headA; p != nil; p, lenA = p.Next, lenA+1 {
	}
	for p := headB; p != nil; p, lenB = p.Next, lenB+1 {
	}
	var long, short *ListNode
	var shortLen int
	var abs int
	if lenA > lenB {
		long, short = headA, headB
		shortLen = lenB
		abs = lenA - lenB
	} else {
		long, short = headB, headA
		shortLen = lenA
		abs = lenB - lenA
	}
	// 第二次遍历, 长链表先走
	for i := 0; i < abs; i++ {
		long = long.Next
	}
	for i := 0; i < shortLen; i++ {
		if long == short {
			return short
		}
		long = long.Next
		short = short.Next
	}
	return short
}

// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m > n {
		return nil
	}
	dummy := &ListNode{Next: head}
	prev := dummy

	//走到将要翻转节点的前一个节点 prev
	for i := 0; i < m-1; i++ {
		prve = prve.Next
	}

	//cur 第m个节点，也就是将要翻转的节点
	cur := prve.Next
	for i := m; i < n; i++ {
		tmp := cur.Next     //保存要反转节点的下一个节点
		cur.Next = tmp.Next //当前节点指向 要放转节点的next节点，最终指向第m个节点的next
		tmp.Next = prev.Next
		prev.Next = tmp
	}
	return dummy.Next
}

// 判断链表是否有环
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 移除链表指定值的元素
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
