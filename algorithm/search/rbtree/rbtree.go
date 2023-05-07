package rbtree

const (
	RED   = false
	BLACK = true
)

type Compare func(a, b interface{}) int

// 红黑树是高级的二叉查找树, 具有二叉查找树的所有性质
type RedBlackTree struct {
	Compare
	root *Node
	size int
}

func (r *RedBlackTree) Size() int {
	return r.size
}

func (r *RedBlackTree) Max() interface{} {
	if r.root == nil {
		return nil
	}
	return r.root.max().value
}

func (r *RedBlackTree) Min() interface{} {
	if r.root == nil {
		return nil
	}
	return r.root.min().value
}

func (r *RedBlackTree) replace(src, dest *Node) {
	if src.parent == nil {
		r.root = dest
	} else {
		if src == src.parent.left {
			src.parent.left = dest
		} else {
			src.parent.right = dest
		}
	}
	if dest != nil {
		dest.parent = src.parent
	}
}

// 左旋
func (r *RedBlackTree) rotateLeft(src *Node) {
	dest := src.right
	r.replace(src, dest)
	src.right = dest.left
	if dest.left != nil {
		dest.left.parent = src
	}
	dest.left = src
	src.parent = dest
}

// 右旋
func (r *RedBlackTree) rotateRight(src *Node) {
	dest := src.left
	r.replace(src, dest)
	src.left = dest.right
	if dest.right != nil {
		dest.right.parent = src
	}
	dest.right = src
	src.parent = dest
}

//////////////////////////////////////////////////////
// 红黑树节点
type Node struct {
	value               interface{}
	left, right, parent *Node
	color               bool
}

// 返回叔叔节点
func (n *Node) sibling() *Node {
	if n == n.parent.left {
		return n.parent.right
	}
	return n.parent.left
}

func (n *Node) max() *Node {
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *Node) min() *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

// 返回节点颜色
func (n *Node) defcolor() bool {
	if n == nil {
		return BLACK
	}
	return n.color
}
