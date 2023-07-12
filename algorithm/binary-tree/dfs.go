package binary_tree

import (
	"fmt"
	"golearn/algorithm/util"
)

// 递归模版 深度优先遍历（前序）
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	// Do something
	fmt.Println(root.Val)
	dfs(root.Left)
	dfs(root.Right)
}

// 左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		ans += root.Left.Val
	}
	if root.Left != nil {
		ans += sumOfLeftLeaves(root.Left)
	}
	if root.Right != nil {
		ans += sumOfLeftLeaves(root.Right)
	}
	return ans
}

// 给定一棵二叉树，以逆时针顺序从根开始返回其边界。边界按顺序包括左边界、叶子结点和右边界而不包括重复的结点。 (结点的值可能重复)
// dfs + 左右边界标记
// 简单理解: 一个先序遍历, 给节点标记flag
// https://leetcode-cn.com/problems/boundary-of-binary-tree/solution/er-cha-shu-de-bian-jie-by-leetcode/
func boundaryOfBinaryTree(root *TreeNode) []int {
	var res = make([]int, 0)
	boundary(root, true, true, &res)
	return res
}

func boundary(node *TreeNode, leftBound, rightBound bool, res *[]int) {
	if node == nil {
		return
	}
	if leftBound {
		*res = append(*res, node.Val)
	} else if node.Left == nil && node.Right == nil {
		*res = append(*res, node.Val)
		return
	}
	boundary(node.Left, leftBound, !leftBound && rightBound && node.Right == nil, res)
	boundary(node.Right, !rightBound && leftBound && node.Left == nil, rightBound, res)
	if !leftBound && rightBound {
		*res = append(*res, node.Val)
	}
}

// 求最大深度
func MaxDepth(head *TreeNode) int {
	if head == nil {
		return 0
	}
	left := MaxDepth(head.Left)
	right := MaxDepth(head.Right)
	//注意退出条件，取左右子树的最大深度，再加上 1（本节点）
	return util.Max(left, right) + 1
}

var maxDepth int
var curDepth int

// 用回溯思想求最大深度
func MaxDepthBackstrack(head *TreeNode) int {
	var traverse = func(head *TreeNode) {
		if head == nil {
			return
		}
		curDepth++
		if head.Left == nil && head.Right == nil {
			maxDepth = util.Max(maxDepth, curDepth)
		}
		MaxDepthBackstrack(head.Left)
		MaxDepthBackstrack(head.Right)
		curDepth--
	}

	traverse(head)
	return maxDepth
}

// 最小深度
func MinDepth(head *TreeNode) int {
	if head == nil {
		return 0
	}
	// null节点不参与比较
	if head.Left != nil && head.Right == nil {
		return 1 + MinDepth(head.Left)
	}
	if head.Left == nil && head.Right != nil {
		return 1 + MinDepth(head.Right)
	}
	return util.Min(MinDepth(head.Left), MinDepth(head.Right)) + 1
}

// 计算二叉树节点数
func Count(root *TreeNode) int {
	var traverse func(*TreeNode) int

	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		var left = traverse(root.Left)
		var right = traverse(root.Right)

		return left + right + 1
	}

	return traverse(root)
}
