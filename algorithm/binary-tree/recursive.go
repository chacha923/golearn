package binary_tree

import "fmt"

// 递归 + 分治 思想
// 分解为子问题，子问题的成功条件和父问题一样
// 子问题的结果聚合后就是父问题的结果

// 镜像二叉树
func MirrorRecursively(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}

	swapLRNode(root)
	if root.Left != nil {
		MirrorRecursively(root.Left)
	}
	if root.Right != nil {
		MirrorRecursively(root.Right)
	}
}

// 左右孩子互换
func swapLRNode(father *TreeNode) {
	tmp := father.Left
	father.Left = father.Right
	father.Right = tmp
}

// 判断二叉树是否对称(镜像)
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsIsSymmetric(root.Left, root.Right)
}

func dfsIsSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left != nil && right == nil) || (left == nil && right != nil) || (left.Val != right.Val) {
		return false
	}
	return dfsIsSymmetric(left.Left, right.Right) && dfsIsSymmetric(left.Right, right.Left)
}

func printArray(array []int, i int) {
	if i == len(array) {
		return
	}
	fmt.Println(array[i])
	printArray(array, i+1)
}
