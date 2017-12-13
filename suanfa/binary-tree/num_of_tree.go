package binary_tree

func NumOfTree(root *TreeNode)int{
	if root == nil {
		return 0
	}
	left := NumOfTree(root.left)
	right := NumOfTree(root.right)
	return left + right + 1
}
