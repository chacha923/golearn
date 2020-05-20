package binarytree

//求二叉树第k层节点数
func numsOfkLevelTreeNode(root *TreeNode, k int) int {
	if root == nil || k < 1 {
		return 0
	}

	if k == 1 { // 根节点
		return 1
	}

	numLeft := numsOfkLevelTreeNode(root.Left, k-1)
	numRight := numsOfkLevelTreeNode(root.Right, k-1)
	return numLeft + numRight
}
