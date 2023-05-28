package search

import (
	binary_tree "golearn/algorithm/binary-tree"
	"golearn/algorithm/list"
)

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

func Search(root *binary_tree.TreeNode, data int) *binary_tree.TreeNode {
	if root == nil {
		return nil
	}
	var current = root
	for current != nil {
		if data < current.Val {
			current = current.Left
		}
		if data > current.Val {
			current = current.Right
		}
		if data == current.Val {
			return current
		}
	}
	return nil
}

func SearchRecurse(root *binary_tree.TreeNode, data int) *binary_tree.TreeNode {
	if root == nil {
		return nil
	}
	if data == root.Val {
		return root
	}
	if data < root.Val {
		return SearchRecurse(root.Left, data)
	}
	if data > root.Val {
		return SearchRecurse(root.Right, data)
	}
	return nil
}

func Delete(root *binary_tree.TreeNode, data int) *binary_tree.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == data {
		return root
	}
}

// 将左子树转换为双向链表。
// 找到左子树双向链表的最后一个节点。
// 如果左子树双向链表的最后一个节点不为空，将其与当前节点链接。
// 将右子树转换为双向链表。
// 如果右子树双向链表的第一个节点不为空，将其与当前节点链接。
// 返回链表的头节点。
// 该算法的时间复杂度为 O(n)
func Bst2DoubleList(root *binary_tree.TreeNode) *binary_tree.TreeNode {
	// 双向链表也可以用 TreeNode 结构体表示节点，left -> prev, right -> next
	// 返回的是双向链表就是完全退化后的二叉查找树
	if root == nil {
		return nil
	}
	// 1. 拿到左子树转链表后的头结点
	var leftHead = Bst2DoubleList(root.Left)
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
	var rightHead = Bst2DoubleList(root.Right)
	// 接到 root 后面
	if rightHead != nil {
		rightHead.Left = root
		root.Right = rightHead
	}
	return leftHead
}

// 时间复杂度：O(n)，其中 n 是链表的长度。
// 空间复杂度：O(n)，这里只计算除了返回答案之外的空间。平衡二叉树的高度为 O(logn)O(logn)，即为递归过程中栈的最大深度，也就是需要的空间。
// 给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。
func SortedListToBST(head *list.DoubleNode) *binary_tree.TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return binary_tree.NewTreeNode(head.Val)
	}

	// 1. 找到中间节点，因为要求高度平衡
	// 分治，求有序链表任意一段的中位节点
	var mid = midOfList(head)
	// 整个树的根节点
	var root = binary_tree.NewTreeNode(mid.Val)
	root.Right = SortedListToBST(mid.Next.Next)
	// 截断
	mid.Next = nil
	// 构建右子树
	root.Left = SortedListToBST(head)
	return root
}

// 给出头节点，求链表的中心节点
func midOfList(head *list.DoubleNode) *list.DoubleNode {
	// 来个快慢指针
	var slow = head
	var fast = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
func IsValidBST(root *binary_tree.TreeNode) bool {
	// 节点的左子树只包含 小于 当前节点的数。
	// 节点的右子树只包含 大于 当前节点的数。
	// 所有左子树和右子树自身必须也是二叉搜索树。

	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	// 左孩子必须小于根节点
	if root.Left != nil {
		if root.Val < root.Left.Val {
			return false
		}
	}
	// 右孩子必须大于根节点
	if root.Right != nil {
		if root.Val > root.Right.Val {
			return false
		}
	}
	return IsValidBST(root.Left) && IsValidBST(root.Right)
}

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
func KthSmallest(root *binary_tree.TreeNode, k int) int {
	// 中序遍历找第 k 个节点
	var orderResult = make([]int, 0)
	var dfs = func(*binary_tree.TreeNode) {}
	dfs = func(root *binary_tree.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		// 中序遍历
		orderResult = append(orderResult, root.Val)
		dfs(root.Right)
	}
	return orderResult[k-1]
}
