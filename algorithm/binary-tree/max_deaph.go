package binary_tree

import (
	"golearn/suanfa/lib"
)

func MaxDepth(head *TreeNode) int {
	if head == nil {
		return 0
	}

	left := MaxDepth(head.Left)
	right := MaxDepth(head.Right)
	return lib.Max(left, right) + 1 //注意退出条件
}
