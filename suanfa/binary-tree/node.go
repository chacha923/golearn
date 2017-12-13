package binary_tree

//二叉树节点

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func NewTreeNode (value int) *TreeNode{
	return &TreeNode{value:value}
}