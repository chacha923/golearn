package binary_tree

//二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Val: value}
}

// n叉树节点
type Node struct {
	Val      int
	Children []*Node
}
