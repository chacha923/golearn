package list

import (
	"golearn/algorithm/structure"
	"sort"
)

// 合并两个有序链表 递归法
// 递归函数定义：合并l1, l2 为开头的两个链表，并返回新的头结点
func MergeLists(l1, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// l1 值小，那么 l1.Next 指向子问题结果
	if l1.Val < l2.Val {
		l1.Next = MergeLists(l1.Next, l2)
		return l1
	} else if l1.Val >= l2.Val {
		// l2 值小，那么 l2.Next 指向子问题结果
		l2.Next = MergeLists(l1, l2.Next)
	}
	return l2
}

// 合并两个有序链表 迭代法（拉链）
func MergeLists2(l1, l2 *Node) *Node {
	dummy := NewEmptyNode()
	// 游标指针，每次选择较小的头结点
	cur := dummy
	for l1 != nil && l2 != nil {
		// l1 l2 谁小，谁就是 cur 指针要去的下一个节点
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else if l1.Val >= l2.Val {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	// 处理剩余节点
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

// 通过两个辅助队列实现
func MergeRange3(l1, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &Node{Val: -1}
	prev := dummy
	var q1 = structure.NewQueue[*Node]()
	var q2 = structure.NewQueue[*Node]()
	for l1 != nil {
		q1.Push(l1)
		l1 = l1.Next
	}
	for l2 != nil {
		q2.Push(l2)
		l2 = l2.Next
	}
	for q1.Len() > 0 && q2.Len() > 0 {
		if q1.Peek().Val < q2.Peek().Val {
			prev.Next = q1.Pop()
		} else if q1.Peek().Val >= q2.Peek().Val {
			prev.Next = q2.Pop()
		}
		prev = prev.Next
	}
	if q1.Len() > 0 {
		prev.Next = q1.Pop()
	} else if q2.Len() > 0 {
		prev.Next = q2.Pop()
	}
	return dummy.Next
}

// 合并 N 个有序链表
func MergeNLists(heads []*Node) *Node {
	if len(heads) == 0 {
		return nil
	}
	if len(heads) == 1 {
		return heads[0]
	}
	var n = len(heads)
	for n > 1 {
		k := (n + 1) / 2
		for i := 0; i < n/2; i++ {
			heads[i] = MergeLists(heads[i], heads[i+k])
		}
		n = k
	}
	return heads[0]
}

// 给出一个区间的集合，请合并所有重叠的区间。
// 输入: [[1,3],[2,6],[8,10],[15,18]]
// 输出: [[1,6],[8,10],[15,18]]
// 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
// https://leetcode-cn.com/problems/merge-intervals/solution/he-bing-qu-jian-by-leetcode-solution/
func MergeRange(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	less := func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	}
	sort.Slice(intervals, less)
	merged := [][]int{}
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		m := merged[len(merged)-1]
		c := intervals[i]
		if c[0] > m[1] {
			merged = append(merged, c)
			continue
		}
		if c[1] > m[1] {
			m[1] = c[1]
		}
	}
	return merged
}

// 两两交换链表中的节点
// 给定 1->2->3->4, 你应该返回 2->1->4->3.
func SwapPairs(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	// 因为 head 会发生移动，用一个 dummy 节点使 next 永远指向第一个节点
	dummy := NewEmptyNode()
	dummy.Next = head
	prevNode := dummy
	for head != nil && head.Next != nil {
		// Nodes to be swapped
		firstNode := head
		secondNode := head.Next

		//swapping the
		prevNode.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		// Reinitializing the head and prevNode for next swap
		prevNode = firstNode
		head = firstNode.Next // jump
	}
	return dummy.Next
}
