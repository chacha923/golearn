package binary_tree

import (
	"golearn/algorithm/util"
)

// 一些简单的练手
// 树基本上都用递归思想

// 计算叶子节点数目
func calculateTreeNodeNumber(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return calculateTreeNodeNumber(root.Left) + calculateTreeNodeNumber(root.Right) + 1
}

// 计算一棵二叉树的最长直径长度。
// 解决这题的关键在于，每一条二叉树的「直径」长度，就是一个节点的左右子树的最大深度之和。
func DiameterOfBinaryTree(root *TreeNode) int {
	// 遍历二叉树节点，计算每个节点的左右子树的最大深度之和，取最大值
	var maxDiameter = 0

	var traverse = func(root *TreeNode) {}
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		var leftDepth = MaxDepth(root.Left)
		var rightDepth = MaxDepth(root.Right)
		maxDiameter = util.Max(maxDiameter, leftDepth+rightDepth)
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return maxDiameter
}
