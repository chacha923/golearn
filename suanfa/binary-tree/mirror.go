package binary_tree

func MirrorRecursively(root *TreeNode) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		return
	}

	swapLRNode(root)
	if root.left != nil {
		MirrorRecursively(root.left)
	}
	if root.right != nil {
		MirrorRecursively(root.right)
	}
}

func swapLRNode(father *TreeNode) {
	tmp := father.left
	father.left = father.right
	father.right = tmp
}
