package binary_tree

func MirrorRecursively(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}

	swapLRNode(root)
	if root.Left != nil {
		MirrorRecursively(root.Left)
	}
	if root.Right != nil {
		MirrorRecursively(root.Right)
	}
}

// 左右孩子互换
func swapLRNode(father *TreeNode) {
	tmp := father.Left
	father.Left = father.Right
	father.Right = tmp
}
