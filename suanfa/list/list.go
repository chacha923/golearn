package list

// 解决链表问题最好的办法是在脑中或者纸上把链表画出来

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

// 判断链表是否有环
// 快指针追上慢指针说明有环
// 快指针走到链表尾说明无环
// https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
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

	// 求环的起始位置
	// 当快慢指针相遇时，让其中任一个指针指向头节点，然后让它俩以相同速度前进，
	// 再次相遇时所在的节点位置就是环开始的位置。
	// slow = head
	// for slow != fast {
	// 	  fast = fast.Next
	// 	  slow = slow.Next
	// }
	// return slow
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

func kthToLast(head *ListNode, k int) int {
	if k <= 0 || head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head
	for i := 0; i < k; i++ {
		dummy = dummy.Next
	}
	for {
		if dummy.Next == nil {
			return head
		}
		dummy = dummy.Next
		head = head.Next
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

// 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mtable := make(map[int]struct{})
	cur := head
	mtable[cur.Val] = struct{}{}
	for cur != nil && cur.Next != nil {
		// 不重复
		if _, ok := mtable[cur.Next.Val]; !ok {
			mtable[cur.Next.Val] = struct{}{}
			cur = cur.Next
			continue
		}
		// 重复了删除, 注意可能有多个连续重复值, 直到遇到一个不重复值cur才往前走
		cur.Next = cur.Next.Next
	}
	return head
}

// 请判断一个链表是否为回文链表。
// 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast := head
	slow := head
	for {
		// 注意退出条件, 奇数长度fast走到倒数1节点, 偶数长度fast走到倒数第2节点
		if fast.Next == nil || fast.Next.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	// slow 后面的部分反转, pre指向反转后的头结点
	// 长度为奇数, slow指向的刚好是中间节点, 因此后面的值比较不带中间节点了
	pre := reverse(slow.Next)
	for head != nil && pre != nil {
		if head.Val != pre.Val {
			return false
		}
		head = head.Next
		pre = pre.Next
	}
	return true
}

// 寻找重复数
// 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
// 数组映射为链表, 转化为寻找链表环的起始节点, 下标视为当前节点val, 元素视为next节点的val
// 如果数组中有重复的数，以数组 [1,3,4,2,2] 为例,我们将数组下标 n 和数 nums[n] 建立一个映射关系 f(n)f(n)，其映射关系 n->f(n)
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// 奇偶链表
// 给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
// 请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。
// 输入: 2->1->3->5->6->4->7->NULL
// 输出: 2->3->6->7->1->5->4->NULL
// 应当保持奇数节点和偶数节点的相对顺序。
// 链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	// head => 奇数组头
	odd := head                            // 奇数组尾
	evenHead, even := head.Next, head.Next // 偶数组头尾
	//odd -> even -> odd -> even -> odd
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// 两数相加
// 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
// 输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 8 -> 0 -> 7
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 用两个栈
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	var carry int //进位
	var head *ListNode
	for len(s1) != 0 || len(s2) != 0 || carry != 0 {
		tmp := 0
		// 同时pop出一位数
		if len(s1) > 0 {
			tmp += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			tmp += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		tmp += carry
		ans := 0
		carry = tmp / 10 // 如果进位 carry == 1, 否则0
		ans = tmp % 10   // 当前位数
		// 从后往前(低位到高位)加节点
		cur := &ListNode{
			Val:  ans,
			Next: head,
		}
		head = cur
	}
	return head
}

// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return []int{head.Val}
	}
	stack := make([]int, 0)
	res := make([]int, 0)

	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top)
	}
	return res
}
