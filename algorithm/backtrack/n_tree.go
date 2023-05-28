package backtrack

import binary_tree "golearn/algorithm/binary-tree"

// 解决 N 叉树问题，但是用遍历回溯的思想

func NTreePreOrder(root *binary_tree.NTreeNode) {
	var traverse func(root *binary_tree.NTreeNode)
	traverse = func(root *binary_tree.NTreeNode) {
		if root == nil {
			return
		}
		// 前序遍历位置, do something
		// res.add(root.val)
		for _, child := range root.Children {
			// 前序遍历 回溯，但是把根节点踢除了
			// do
			traverse(child)
			// 回溯
		}
		// 后续遍历位置, do something
		// res.add(root.val)
	}

	traverse(root)
}
