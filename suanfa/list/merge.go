package list

import "sort"

// 合并两个有序链表 递归法
func MergeLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeLists(l1.Next, l2)
		return l1
	}
	l2.Next = MergeLists(l1, l2.Next)
	return l2
}

// 合并两个有序链表 迭代法
func MergeLists2(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}
	return dummy.Next
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
func SwapPairs(head *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := NewListNode(-1)
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
