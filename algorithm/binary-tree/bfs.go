package binary_tree

import "golearn/algorithm/structure"

// https://www.zhihu.com/question/41911660 遍历二叉树深度优先和广度优先的通俗解释？

// dfs：进栈、退栈，一搜到底
// bfs：入队、出队，步步为营

// 深度优先，就是一条路走到底，广度优先，就是每条路都同时派人走。
// 这都是遍历一个集合的方式，目的一样，只是实现方式不同。
// 我喜欢深度优先，因为可以使用递归，简单。但广度优先很容易使用并发算法优化。
// 如果有一种极端情况，这个结构很大，使用递归会耗尽函数的栈空间。
// 这个时候，深度优先算法就不能用了，就要考虑使用广度优先算法（层次遍历）。

// 二叉树最大宽度
// 给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。
// 每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
func WidthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 1
	// queue := []Item{{0, root}} // bfs使用的队列
	var queue = structure.NewQueue[*TreeNode]()
	queue.Push(root)
	for queue.Len() > 0 {
		// 当前层的宽度
		l := queue.Len()
		if l > ans {
			ans = l
		}
		// 每一次循环, queue内包含同一层的节点(分层遍历?)
		for i := 0; i < queue.Len(); i++ {
			var pop = queue.Pop()
			if pop.Left != nil {
				queue.Push(pop.Left)
			}
			if pop.Right != nil {
				queue.Push(pop.Right)
			}
		}
	}
	return ans
}

// 返回二叉树的第k层节点
// 同样是 bfs 思想，就是每次记录层数，记录当前队列长度，卡住每次循环的出队次数
func KthLevel(root *TreeNode, k int) []int {
	if root == nil || k < 1 {
		return nil
	}
	var currentLevel = 1
	var result = make([]int, 0)
	var queue = structure.NewQueue[*TreeNode]()
	queue.Push(root)
	for queue.Len() > 0 {
		var currentSize = queue.Len()
		for i := 0; i < currentSize; i++ {
			var node = queue.Pop()
			if currentLevel == k {
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		currentLevel++
	}
	return result
}
