package search

import binary_tree "golearn/algorithm/binary-tree"

// 二叉搜索树
type BST struct {
	// 保存根节点，用于搜索
	root *binary_tree.TreeNode
}

func (cls *BST) Insert(data int) {
	var newNode = binary_tree.NewTreeNode(data)
	// 没有根，自己做根
	if cls.root == nil {
		cls.root = newNode
		return
	}
	var current = cls.root
	// 新节点应该插到哪个位置
	for {
		if data < current.Val {
			// 往左走 找空位
			if current.Left == nil {
				current.Left = newNode
				return
			}
			// 没空位 继续往左
			current = current.Left
		} else if data > current.Val {
			// 往右走
			if current.Right == nil {
				current.Right = newNode
				return
			}
			// 没空位，继续往右
			current = current.Right
		}
	}
}

func (cls *BST) Search(data int) bool {
	if cls.root == nil {
		return false
	}
	var current = cls.root
	for current != nil {
		if data < current.Val {
			current = current.Left
		}
		if data > current.Val {
			current = current.Right
		}
		if data == current.Val {
			return true
		}
	}
	return false
}

// 将左子树转换为双向链表。
// 找到左子树双向链表的最后一个节点。
// 如果左子树双向链表的最后一个节点不为空，将其与当前节点链接。
// 将右子树转换为双向链表。
// 如果右子树双向链表的第一个节点不为空，将其与当前节点链接。
// 返回链表的头节点。
// 该算法的时间复杂度为 O(n)
func bst2DoubleList(root *binary_tree.TreeNode) *binary_tree.TreeNode {
	// 双向链表也可以用 TreeNode 结构体表示节点，left -> prev, right -> next
	// 返回的是双向链表就是完全退化后的二叉查找树
	if root == nil {
		return nil
	}
	// 1. 拿到左子树转链表后的头结点
	var leftHead = bst2DoubleList(root.Left)
	// 拿到尾节点，也就是无限往右找
	var leftTail = leftHead
	for leftTail != nil && leftTail.Right != nil {
		leftTail = leftTail.Right
	}

	// 根接到左链表后面
	if leftTail != nil {
		leftTail.Right = root
		root.Left = leftTail
	}

	// 操作右子树
	// 拿到右子树转换后的头结点
	var rightHead = bst2DoubleList(root.Right)
	// 接到 root 后面
	if rightHead != nil {
		rightHead.Left = root
		root.Right = rightHead
	}
	return leftHead
}
