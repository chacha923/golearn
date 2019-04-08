package binary_tree

func NumOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := NumOfTree(root.Left)
	right := NumOfTree(root.Right)
	return left + right + 1
}
