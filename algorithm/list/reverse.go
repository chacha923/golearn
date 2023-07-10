package list

import "golearn/algorithm/structure"

// 反转链表节点，只需要记住《前序节点》，链表就不会断
// temp		temp.next  temp.next.next
// temp	->	node1  -> node2

// 反转链表，返回新的头结点（原来的最后一个节点）
// 容易想不到，不要上来就写这种
func ReverseList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	var prev = head
	// head 先走一步，这里用个 dummy 节点更优雅
	head = head.Next
	// 释放第一个节点的指针域
	prev.Next = nil
	// 1 (prev) ->  2 (root) ->  3 (root.Next)
	for head != nil {
		var tmp = head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return head
}

// 递归 不用辅助栈
func ReverseList1(head *Node) *Node {
	// 定义：输入一个单链表头结点，将该链表反转，返回新的头结点
	var reverse func(head *Node) *Node
	reverse = func(head *Node) *Node {
		// 最多递归到倒数第二个节点
		if head == nil || head.Next == nil {
			return head
		}
		var last = reverse(head.Next)
		// 反转
		head.Next.Next = head
		// 可以直接断开, 因为递归栈保存了前面的节点
		head.Next = nil
		return last
	}
	// last 肯定指向原链表最后一个节点不会变
	// 理解：递归到最后一层时，这个时候执行 return 了，递归栈后面 return 的值就拿不到了。只能拿栈顶的 return 值？
	// 这是符合递归函数定义的，输入 head 节点，返回这个链表反转后的头结点
	return reverse(head)
}

// 辅助栈
func ReverseList2(root *Node) *Node {
	if root == nil || root.Next == nil {
		return root
	}
	var stack = structure.NewStack[*Node]()
	for root != nil {
		stack.Push(root)
		root = root.Next
	}
	var dummy = NewEmptyNode()
	dummy.Next = stack.Top()
	for stack.Len() > 0 {
		var pop = stack.Pop()
		if stack.Len() > 0 {
			//  指向栈顶节点
			pop.Next = stack.Top()
		} else if stack.Len() == 0 {
			// 栈空，说明 pop 是最后一个节点
			pop.Next = nil
		}
	}
	return dummy.Next
}

// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL
func reverseBetween(head *Node, m int, n int) *Node {
	if head == nil || m > n {
		return nil
	}
	dummy := &Node{Next: head}
	prev := dummy

	//走到将要翻转节点的前一个节点 prev
	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}

	//cur 第m个节点，也就是将要翻转的节点
	cur := prev.Next
	for i := m; i < n; i++ {
		tmp := cur.Next     //保存要反转节点的下一个节点
		cur.Next = tmp.Next //当前节点指向 要放转节点的next节点，最终指向第m个节点的next
		tmp.Next = prev.Next
		prev.Next = tmp
	}
	return dummy.Next
}
