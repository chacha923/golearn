package binarytree

// 二叉树节点数
func NumOfTreeNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := NumOfTreeNode(root.Left)
	right := NumOfTreeNode(root.Right)
	return left + right + 1
}
