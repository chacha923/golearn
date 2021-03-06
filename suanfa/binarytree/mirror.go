package binarytree

// 镜像二叉树
func MirrorRecursively(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}

	swapLRNode(root)
	MirrorRecursively(root.Left)
	MirrorRecursively(root.Right)
}

// 交换左右孩子值
func swapLRNode(father *TreeNode) {
	tmp := father.Left
	father.Left = father.Right
	father.Right = tmp
}
