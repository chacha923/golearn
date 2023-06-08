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
func reverseList(root *Node) *Node {
	if root == nil || root.Next == nil {
		return root
	}
	var prev = root
	root = root.Next
	// 1 (prev) ->  2 (root) ->  3 (root.Next)
	for root != nil {
		var tmp = root.Next
		root.Next = prev
		prev = root
		root = tmp
	}
	return root
}

// 递归 不用辅助栈
func ReverseList1(root *Node) *Node {
	if root == nil || root.Next == nil {
		return nil
	}
	// 最后找不到头，需要提前保存头节点
	var head = root
	for head.Next != nil {
		head = head.Next
	}
	var reverse = func(root *Node) {}
	reverse = func(root *Node) {
		// 最多递归到倒数第二个节点
		if root.Next.Next == nil {
			return
		}
		reverse(root.Next)
		// 反转
		root.Next.Next = root
	}
	return head
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
