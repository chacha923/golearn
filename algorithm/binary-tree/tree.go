package binary_tree

import (
	"golearn/algorithm/structure"
	"golearn/algorithm/util"
)

var maxLen int = 0
var maxLenPath []int = make([]int, 0)
var stack = structure.NewStack[int]()

// 打印二叉树的最长路径
func FindLongestPath(root *TreeNode) []int {
	// 考虑用栈？
	// 1. 不断做深度遍历，一旦到底了，开始考察栈中的元素

	dfsAndPutStack(root, stack)
	return maxLenPath
}

func dfsAndPutStack(root *TreeNode, stack *structure.Stack[int]) {
	if root == nil {
		// 到底了，开始检查栈
		if stack.Len() > maxLen {
			maxLen = stack.Len()
			maxLenPath = stack.ToSlice()
		}
		return
	}
	// 否则进栈
	stack.Push(root.Val)
	dfsAndPutStack(root.Left, stack)
	dfsAndPutStack(root.Right, stack)
	// 左右子树都找完了，出栈
	stack.Pop()
}

// 不用辅助栈实现，就是把求最大深度的代码改一下
// 后续遍历
func FindLongestPath1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	leftPath := FindLongestPath1(root.Left)
	rightPath := FindLongestPath1(root.Right)

	if len(leftPath) >= len(rightPath) {
		return append(leftPath, root.Val)
	}
	return append(rightPath, root.Val)
}

// 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
func FindBottomLeftValue(root *TreeNode) int {
	// 层次遍历，但是每层从右往左，那么队列的最后一个节点就是目标节点
	var nodeQueue = structure.NewQueue[*TreeNode]()
	nodeQueue.Push(root)
	for nodeQueue.Len() > 0 {
		var pop = nodeQueue.Pop()
		if pop.Right != nil {
			nodeQueue.Push(pop.Right)
		}
		if pop.Left != nil {
			nodeQueue.Push(pop.Left)
		}
	}
	return root.Val
}

// 判断一颗二叉树树是不是平衡二叉树
func IsBalance(root *TreeNode, depth *int) bool {
	if root == nil {
		return true
	}
	left := 0
	right := 0
	if IsBalance(root.Left, &left) && IsBalance(root.Right, &right) {
		diff := left - right
		if diff <= 1 && diff >= -1 {
			return true
		}
	}
	return false
}

// 利用求二叉树深度, 效率低
func IsBalance1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	nLeftDepth := MaxDepth(root.Left)
	nRightDepth := MaxDepth(root.Right)
	diff := nRightDepth - nLeftDepth

	if diff > 1 || diff < -1 {
		return false
	}
	return IsBalance1(root.Left) && IsBalance1(root.Right)
}

// 给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k 。
// val 所有值 不同
// 返回到目标结点 target 距离为 k 的所有结点的值的列表。 答案可以以 任何顺序 返回。
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	// 用到图的思想，树是特殊的有向图，但我们要用 hashmap 记录父节点，就能同时支持向上搜索
	// 或者改造树的结构，增加一个指向父节点的指针，那么就是 3 条路可以走。（直接转为图的bfs也可以）
	if root == nil {
		return nil
	}
	var graph = make(map[int][]int)
	var generateGraphFromTree func(root *TreeNode)
	generateGraphFromTree = func(root *TreeNode) {
		if root.Left != nil {
			// 双向边
			graph[root.Val] = append(graph[root.Val], root.Left.Val)
			graph[root.Left.Val] = append(graph[root.Left.Val], root.Val)
			generateGraphFromTree(root.Left)
		}
		if root.Right != nil {
			graph[root.Val] = append(graph[root.Val], root.Right.Val)
			graph[root.Right.Val] = append(graph[root.Right.Val], root.Val)
			generateGraphFromTree(root.Right)
		}
	}
	var targetVal = target.Val
	var queue = structure.NewQueue[int]()
	var result = make([]int, 0)
	queue.Push(targetVal)
	for i := 0; i < k-1; i++ {
		// 队列中的元素都是距离 target 距离为 i 的节点
		currentSize := queue.Len()
		for j := 0; j < currentSize; j++ {
			inst := queue.Pop()
			childOfInst := graph[inst]
			for _, child := range childOfInst {
				queue.Push(child)
			}
		}
	}
	for queue.Len() > 0 {
		result = append(result, queue.Pop())
	}
	// 结果要去重
	return util.RemoveDuplicates(result)
}
