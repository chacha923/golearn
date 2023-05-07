package search

import binary_tree "golearn/algorithm/binary-tree"

//二分查找，返回key值在数组中的下标，否则返回-1
//要求数组有序
//时间复杂度 logn
func binarySearch(array []int, key int) int {
	var left = 0
	var right = len(array) - 1

	for left <= right {
		// Prevent (left + right) overflow
		// var mid = (left + right) / 2
		var mid = left + (right-left)/2 // 防止溢出
		if array[mid] == key {
			return mid
		} else if array[mid] < key {
			left = mid + 1
		} else if array[mid] > key {
			right = mid - 1
		}
	}
	// End Condition: left > right
	return -1
}

// 递归的二分查找

func binarySearch2(array []int, key int) int {
	return binarySearchRecursive(array, 0, len(array)-1, key)
}

// left 左下标，right 右下标，将在 array 的左右下标范围内查找
func binarySearchRecursive(array []int, left, right int, target int) int {
	// 边界条件
	if left > right {
		return -1
	}
	var mid = left + (right-left)/2 // 防止溢出
	if array[mid] == target {
		return mid
	}
	if array[mid] < target {
		return binarySearchRecursive(array, mid+1, right, target)
	}
	//  array[mid] > target
	return binarySearchRecursive(array, left, mid-1, target)
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
