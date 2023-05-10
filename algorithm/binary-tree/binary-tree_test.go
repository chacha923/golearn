package binary_tree

import (
	"fmt"
	"testing"
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
	head.Val = 100
	head.Left = NewTreeNode(16)
	head.Left.Left = NewTreeNode(4)
	head.Left.Right = NewTreeNode(8)
	head.Right = NewTreeNode(70)
	head.Right.Left = NewTreeNode(2)
	head.Right.Left.Left = NewTreeNode(23)
	head.Right.Left.Right = NewTreeNode(5)
	head.Right.Right = NewTreeNode(37)
	head.Right.Right.Left = NewTreeNode(12)
}

func TestMaxDeaph(t *testing.T) {
	fmt.Println("max deaph is ", MaxDepth(head))
}

func TestMostLongPath(t *testing.T) {
	var res = FindLongestPath(head)
	fmt.Print(res)
}

func TestFindLongestPath(t *testing.T) {
	var res = FindLongestPath1(head)
	fmt.Print(res)
}
