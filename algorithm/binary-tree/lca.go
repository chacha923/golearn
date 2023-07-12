package binary_tree

// 公共祖先

// 最简单的情况，假设 p q 节点一定存在
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return find(root, p, q)
}

// 如果一个节点能够在它的左右子树中分别找到 p 和 q，则该节点为 LCA 节点。
func find(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 前序
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	var left = find(root.Left, p, q)
	var right = find(root.Right, p, q)
	// 后序位置，已经知道左右子树是否存在目标值
	if left != nil && right != nil {
		// 当前节点是 LCA 节点
		return root
	}

	if left != nil {
		return left
	}
	return right
}
