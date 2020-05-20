package binarytree

import "golearn/suanfa/lib"

//判断一颗二叉树树是不是平衡二叉树
func IsBalance(root *TreeNode, depth *int) bool {
	if root == nil {
		return true
	}
	left := 0
	right := 0
	if IsBalance(root.Left, &left) && IsBalance(root.Right, &right) {
		diff := left - right
		if diff <= 1 && diff >= -1 {
			return true
		}
	}

	return false
}

//利用求二叉树深度, 效率低
func IsBalance1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	nLeftDepth := MaxDepth(root.Left)
	nRightDepth := MaxDepth(root.Right)
	diff := nRightDepth - nLeftDepth

	if diff > 1 || diff < -1 {
		return false
	}
	return IsBalance1(root.Left) && IsBalance1(root.Right)
}

// 求最大深度
func MaxDepth(head *TreeNode) int {
	if head == nil {
		return 0
	}
	left := MaxDepth(head.Left)
	right := MaxDepth(head.Right)
	return lib.Max(left, right) + 1 //注意退出条件
}
