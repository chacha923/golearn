package binary_tree

import (
	"golearn/algorithm/structure"
)

var maxLen int = 0
var maxLenPath []int = make([]int, 0)
var stack = structure.NewStack()

// 打印二叉树的最长路径
func FindLongestPath(root *TreeNode) []int {
	// 考虑用栈？
	// 1. 不断做深度遍历，一旦到底了，开始考察栈中的元素

	dfsAndPutStack(root, stack)
	return maxLenPath
}

func dfsAndPutStack(root *TreeNode, stack *structure.Stack) {
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
