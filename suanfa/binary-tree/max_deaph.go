package binary_tree

import (
	"golearn/suanfa/lib"
)

func MaxDeaph(head *TreeNode) int{
	if head == nil {
		return 0
	}

	left := MaxDeaph(head.left)
	right := MaxDeaph(head.right)
	return lib.Max(left, right) + 1
}
