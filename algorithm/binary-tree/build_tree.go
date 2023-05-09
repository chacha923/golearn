package binary_tree

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
