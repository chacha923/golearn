package backtrack

import (
	binary_tree "golearn/algorithm/binary-tree"
	"golearn/algorithm/util"
	"math"
)

// 解决二叉树问题，但是用遍历回溯的思想

// 翻转二叉树
func InvertTree(root *binary_tree.TreeNode) *binary_tree.TreeNode {
	// 先定义变量，匿名函数才能递归调用
	var traverse func(root *binary_tree.TreeNode)

	traverse = func(root *binary_tree.TreeNode) {
		if root == nil {
			return
		}
		// 交换左右子树
		var temp = root.Left
		root.Left = root.Right
		root.Right = temp

		traverse(root.Left)
		traverse(root.Right)
	}

	traverse(root)
	return root
}

func MaxDepth(root *binary_tree.TreeNode) int {
	var res = 0   // 记录最大深度
	var depth = 0 // 记录当前深度
	var traverse func(root *binary_tree.TreeNode)
	traverse = func(root *binary_tree.TreeNode) {
		if root == nil {
			return
		}
		depth++
		res = util.Max(res, depth)
		traverse(root.Left)
		traverse(root.Right)
		depth--
	}
	return res
}

func MinDepth(root *binary_tree.TreeNode) int {
	var res = math.MaxInt
	var depth = 0
	var traverse func(root *binary_tree.TreeNode)

	traverse = func(root *binary_tree.TreeNode) {
		if root == nil {
			return
		}
		depth++
		if root.IsLeaf()  {
			res = util.Min(res, depth)
		}
		traverse(root.Left)
		traverse(root.Right)
		depth--
	}
	return res
}
