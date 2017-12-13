package binary_tree

import (
	"testing"
	"fmt"
)

//var arr = []int{100, 16, 4, 8, 70, 2, 37, 23, 5, 12}
/*
		  100
	   /       \
	  16        70
	 /  \     /   \
	4    8   2     37
			/  \   /
		   23   5 12
*/
var head *TreeNode
func init() {
	 head = new(TreeNode)
	 head.value = 100
	 head.left = NewTreeNode(16)
	 head.left.left = NewTreeNode(4)
	 head.left.right = NewTreeNode(8)
	 head.right = NewTreeNode(70)
	 head.right.left = NewTreeNode(2)
	 head.right.left.left = NewTreeNode(23)
	 head.right.left.right = NewTreeNode(5)
	 head.right.right = NewTreeNode(37)
	 head.right.right.left = NewTreeNode(12)
}

func TestMaxDeaph(t *testing.T) {
	fmt.Println("max deaph is ",MaxDeaph(head))
}
