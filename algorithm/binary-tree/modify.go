package binary_tree

import (
	"strconv"
	"strings"
)

// 对二叉树做修改

// 删除值为 target 的叶子节点，返回新的根节点
func RemoveLeafNode(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = RemoveLeafNode(root.Left, target)
	root.Right = RemoveLeafNode(root.Right, target)
	// 后序遍历，遇到叶子节点
	if root.Left == nil && root.Right == nil && root.Val == target {
		// 删除就是返回 null
		return nil
	}
	return root
}

// root2 合并到 root1，同位置节点值相加，返回新二叉树的根节点
func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = MergeTrees(root1.Left, root2.Left)
	root1.Right = MergeTrees(root1.Right, root2.Right)
	return root1
}

// 输入某二叉树的`前序遍历和中序遍历`的结果，请重建该二叉树。
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
//
/*	3
   / \
  9   20
	 /  \
    15   7
*/
// 前序遍历特点： 节点按照 [ 根节点 | 左子树 | 右子树 ] 排序，以题目示例为例：[ 3 | 9 | 20 15 7 ]
// 中序遍历特点： 节点按照 [ 左子树 | 根节点 | 右子树 ] 排序，以题目示例为例：[ 9 | 3 | 15 20 7 ]
// 后序遍历特点： 节点按照 [ 左子树 | 右子树 | 根节点 ] 排序，以题目示例为例：[ 9 | 15 7  20 | 3 ]
// 运用到一个特性，中序遍历的根节点左边是左子树，右边是右子树，前序遍历的第一个节点是根节点，后序遍历的最后一个节点是根节点
// 确定了中序遍历的根节点，就可以确定左右子树的节点数，可以用下标操作切割数组，同一个子树的节点在preorder postorder数组内总是连续的
func buildTree1(preOrder []int, inOrder []int) *TreeNode {
	// 构造一个递归调用的辅助函数，用下标，不要原地改入参
	var build func(preOrder []int, preStart int, preEnd int, inOrder []int, inStart int, inEnd int) *TreeNode
	build = func(preOrder []int, preStart int, preEnd int, inOrder []int, inStart int, inEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}
		// 前序遍历的第一个元素就是根节点
		var rootVal = preOrder[preStart]
		var root = NewTreeNode(rootVal)
		var inOrderIndexOfRoot int
		for i := inStart; i <= inEnd; i++ {
			if inOrder[i] == rootVal {
				inOrderIndexOfRoot = i
				break
			}
		}
		var leftSizeOfInOrder = inOrderIndexOfRoot - inStart
		// 接住左右子树
		root.Left = build(preOrder, preStart+1, leftSizeOfInOrder, inOrder, inStart, inOrderIndexOfRoot-1)
		root.Right = build(preOrder, preStart+leftSizeOfInOrder+1, preEnd, inOrder, inOrderIndexOfRoot+1, inEnd)
		return root
	}

	return build(preOrder, 0, len(preOrder)-1, inOrder, 0, len(inOrder)-1)
}

// 根据一棵树的`中序遍历与后序遍历`构造二叉树。
// 中序遍历将二叉树分成左右两棵子树 （左 根 右）
// 后序遍历最后访问根结点 （左 右 根）
// 代码框架和上面一样，只是取根节点的位置不一样
func buildTree2(inorder []int, postorder []int) *TreeNode {
	var build func(inorder []int, inStart int, inEnd int, postOrder []int, postStart int, postEnd int) *TreeNode
	build = func(inorder []int, inStart int, inEnd int, postOrder []int, postStart int, postEnd int) *TreeNode {
		if inStart > inEnd {
			return nil // 越界
		}
		// 后序遍历的最后一个元素就是根节点
		var rootVal = postOrder[postEnd]
		var inOrderIndexOfRoot int
		for i := inStart; i <= inEnd; i++ {
			if inorder[i] == rootVal {
				inOrderIndexOfRoot = i
				break
			}
		}
		var leftSizeOfInOrder = inOrderIndexOfRoot - inStart
		var root = NewTreeNode(rootVal)

		root.Left = build(inorder, inStart, inOrderIndexOfRoot-1, postOrder, postStart, postStart+leftSizeOfInOrder-1)
		root.Right = build(inorder, inOrderIndexOfRoot+1, inEnd, postOrder, postStart+leftSizeOfInOrder, postEnd-1)
		return root
	}

	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

const (
	SEP  = ","
	NULL = "#"
)

// 序列化二叉树（先序遍历）
func Serialize(root *TreeNode, buf *strings.Builder) {
	if root == nil {
		buf.WriteString(NULL)
		buf.WriteString(SEP)
		return
	}
	buf.WriteString(strconv.Itoa(root.Val))
	buf.WriteString(SEP)
	Serialize(root.Left, buf)
	Serialize(root.Right, buf)
}

// 反序列化二叉树（先序遍历）
func Deserialize(data string) *TreeNode {
	// str to list
	var list = strings.Split(data, SEP)

	// 辅助函数
	var deserialize func(list []string) *TreeNode
	deserialize = func(list []string) *TreeNode {
		if len(list) == 0 {
			return nil
		}
		// 出队，第一个字符是根节点
		firstChar := list[0]
		list = list[1:]
		if firstChar == NULL {
			return nil
		}
		// 自顶向下 使用先序遍历构造
		val, _ := strconv.Atoi(firstChar)
		node := NewTreeNode(val)

		node.Left = deserialize(list)
		node.Right = deserialize(list)

		return node
	}

	return deserialize(list)
}
