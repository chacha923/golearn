package binary_tree

func MirrorRecursively(root *TreeNode){
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		return
	}

	tmp := root.left
	root.left = root.right
	root.right = tmp

	if root.left != nil {
		MirrorRecursively(root.left)
	}
	if root.right != nil {
		MirrorRecursively(root.right)
	}
}