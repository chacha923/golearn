package binary_tree

// 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Val: value}
}

func (cls *TreeNode) GetLeft() *TreeNode {
	if cls == nil {
		return nil
	}
	return cls.Left
}

func (cls *TreeNode) GetRight() *TreeNode {
	if cls == nil {
		return nil
	}
	return cls.Right
}

func (cls *TreeNode) SetLeft(val int) {
	if cls == nil {
		return
	}
	cls.Left = NewTreeNode(val)
}

func (cls *TreeNode) SetRight(val int) {
	if cls == nil {
		return
	}
	cls.Right = NewTreeNode(val)
}

// n叉树节点
type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}
