package binarytree

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉树中。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 找不到（nil），或找到了（左子节点或右子节点）
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// 左子树找
	l := lowestCommonAncestor(root.Left, p, q)
	// 右子树找
	r := lowestCommonAncestor(root.Right, p, q)

	if l == nil {
		return r
	} else if r == nil {
		return l
	}
	// 左右都找到了，那说明我就是公共祖先喽
	return root
}

// 二叉树的右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	q := make([]*TreeNode, 0) // bfs队列
	q = append(q, root)

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			// 取最后一个加入到结果集
			v := q[0]
			q = q[1:]
			if i == size-1 {
				ans = append(ans, v.Val)
			}
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
	}
	return ans
}

// 二叉树右视图
// 深度遍历
func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := new([]int)
	dfsWithDeep(root, 0, res) // 深度从0开始计
	return *res
}

func dfsWithDeep(root *TreeNode, deep int, res *[]int) {
	if root == nil {
		return
	}
	// 把每一层第一个元素添加到结果中
	if len(*res) == deep {
		*res = append(*res, root.Val)
	}
	dfsWithDeep(root.Right, deep+1, res)
	dfsWithDeep(root.Left, deep+1, res)
}

//输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。
//假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
//  3
// / \
// 9  20
//  /  \
// 15   7
// 前序遍历特点： 节点按照 [ 根节点 | 左子树 | 右子树 ] 排序，以题目示例为例：[ 3 | 9 | 20 15 7 ]
// 中序遍历特点： 节点按照 [ 左子树 | 根节点 | 右子树 ] 排序，以题目示例为例：[ 9 | 3 | 15 20 7 ]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	var index int
	// 在中序遍历找root
	for i := range inorder {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}
	// 递归操作左右子树
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}
