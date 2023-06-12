package list

import "golearn/algorithm/structure"

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

// 反转链表，返回新的头结点（原来的最后一个节点）
// 容易想不到，不要上来就写这种
func reverseList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	var prev = head
	head = head.Next
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
	var reverse func(head *Node) *Node
	reverse = func(head *Node) *Node {
		// 最多递归到倒数第二个节点
		if head == nil || head.Next == nil {
			return head
		}
		var last = reverse(head.Next)
		// 反转
		head.Next.Next = head
		head.Next = nil
		return last
	}

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
		pop.Next = stack.Top()
	}
	return dummy.Next
}
