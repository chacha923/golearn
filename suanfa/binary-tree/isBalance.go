package binary_tree

//判断一颗二叉树树是不是平衡二叉树

func IsBalance(root *TreeNode, depth *int) bool {
	if root == nil {
		return true
	}
	left := 0
	right := 0
	if IsBalance(root.left, &left) && IsBalance(root.right, &right) {
		diff := left - right
		if diff <= 1 && diff >= -1 {
			return true
		}
	}

	return false
}

//利用求二叉树深度, 效率低
func IsBalance1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nLeftDepth := MaxDeaph(root.left)
	nRightDepth := MaxDeaph(root.right)
	diff := nRightDepth - nLeftDepth

	if diff > 1 || diff < -1 {
		return false
	}
	return IsBalance1(root.left) && IsBalance1(root.right)
}
