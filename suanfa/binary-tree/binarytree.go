package binary_tree

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
// 层次遍历
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
			v := q[0]
			q = q[1:]
			// 取最后一个加入到结果集
			if i == size-1 {
				ans = append(ans, v.Val)
			}
			// 左右孩子进队列
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
	// 取出根节点
	root := &TreeNode{Val: preorder[0]}
	var index int // 保存根节点在中序遍历数组的下标
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

// 给定一个二叉树，原地将它展开为一个单链表。靠右
// 	  1
//   / \
//  2   5
// / \   \
// 3   4   6
// 展开
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6
// 利用后序遍历
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	//将右子树挂到 左子树的最右边
	//再将整个左子树挂到根节点的右边
	if root.Left != nil {
		pre := root.Left // 左子树
		for pre.Right != nil {
			pre = pre.Right // 左子树最右节点
		}
		pre.Right = root.Right // 右子树拼到左子树最右边
		root.Right = root.Left // 左子树挂到右边
		root.Left = nil        // 左边置为空
	}
	root = root.Right
	return
}

// 反向前序遍历
var pre *TreeNode

func flatten1(root *TreeNode) {
	pre = nil
	helper(root)
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Right)
	helper(root.Left)
	//右节点-左节点-根节点 这种顺序正好跟前序遍历相反
	//用pre节点作为媒介，将遍历到的节点前后串联起来
	root.Left = nil
	root.Right = pre
	pre = root
}

// 给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}
	return false
}
