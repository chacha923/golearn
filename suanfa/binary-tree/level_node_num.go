package binary_tree

//求二叉树第k层节点数
func numsOfkLevelTreeNode(root *TreeNode, k int) int{
	if root == nil || k < 1 {
		return 0
	}

	if k == 1 {
		return 1
	}

	numLeft := numsOfkLevelTreeNode(root.left, k-1)
	numRight := numsOfkLevelTreeNode(root.right, k-1)
	return numLeft + numRight
}
