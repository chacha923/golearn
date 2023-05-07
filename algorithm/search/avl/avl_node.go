package avl

type AVLNode struct {
	Val    int
	height int
	Left   *AVLNode
	Right  *AVLNode
}

// 计算当前节点的平衡因子
func (cls *AVLNode) balanceFactor() int {
	if cls == nil {
		return 0
	}
	return cls.Left.balanceFactor() - cls.Right.balanceFactor()
}
