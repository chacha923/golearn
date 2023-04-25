package binary_tree

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

// 根据一棵树的中序遍历与后序遍历构造二叉树。
// 中序遍历将二叉树分成左右两棵子树 （左 根 右）
// 后序遍历最后访问根结点 （左 右 根）
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			return &TreeNode{
				Val:   inorder[i],
				Left:  buildTree(inorder[:i], postorder[:i]),
				Right: buildTree(inorder[i+1:], postorder[i:len(postorder)-1]),
			}
		}
	}
	return nil
}
