package binary_tree

import "golearn/algorithm/util"

// 一些简单的练手
// 树基本上都用递归思想

// 计算叶子节点数目
func calculateTreeNodeNumber(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return calculateTreeNodeNumber(root.Left) + calculateTreeNodeNumber(root.Right) + 1
}

// 计算二叉树的深度
func getTreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftDepth = getTreeDepth(root.Left)
	var rightDepth = getTreeDepth(root.Right)

	return util.Max(leftDepth, rightDepth) + 1
}
