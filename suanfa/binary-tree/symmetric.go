package binary_tree

// 判断二叉树是否对称(镜像)
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsIsSymmetric(root.Left, root.Right)
}

func dfsIsSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left != nil && right == nil) || (left == nil && right != nil) || (left.Val != right.Val) {
		return false
	}
	return dfsIsSymmetric(left.Left, right.Right) && dfsIsSymmetric(left.Right, right.Left)
}
