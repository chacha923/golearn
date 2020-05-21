package binarytree

// 根据前序遍历, 中序遍历构建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if (len(preorder) == 0) || (len(inorder) == 0) {
		return nil
	}
	var rootIdx int
	for i := 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] { // 在中序遍历序列内找到根节点
			rootIdx = i
			break
		}
	}
	// 左右子树归类
	// pre_left, pre_right := preorder[1: rootIdx+1], preorder[rootIdx+1:]
	// in_left, in_right := inorder[0: rootIdx], inorder[rootIdx+1:]
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:rootIdx+1], inorder[0:rootIdx]),
		Right: buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:]),
	}
}
