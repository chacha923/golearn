package backtrack

import (
	"fmt"
	binary_tree "golearn/algorithm/binary-tree"
	"testing"
)

/*
			  100
		   /       \
		  16        70
		 /  \     /   \

	    4    8   2     37

			/  \   /
		   23   5 12
*/
var head *binary_tree.TreeNode

func init() {
	head = new(binary_tree.TreeNode)
	head.Val = 100
	head.Left = binary_tree.NewTreeNode(16)
	head.Left.Left = binary_tree.NewTreeNode(4)
	head.Left.Right = binary_tree.NewTreeNode(8)
	head.Right = binary_tree.NewTreeNode(70)
	head.Right.Left = binary_tree.NewTreeNode(2)
	head.Right.Left.Left = binary_tree.NewTreeNode(23)
	head.Right.Left.Right = binary_tree.NewTreeNode(5)
	head.Right.Right = binary_tree.NewTreeNode(37)
	head.Right.Right.Left = binary_tree.NewTreeNode(12)
}

func TestInvert(t *testing.T) {
	var res = InvertTree(head)
	fmt.Println(res.Right)
}

func TestFullPermute(t *testing.T) {
	var res = FullPermute([]int{1, 2, 3})
	fmt.Println(res)
}
