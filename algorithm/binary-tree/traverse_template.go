package binary_tree

import (
	"fmt"
	"golearn/algorithm/sort"
	"golearn/algorithm/structure"
)

// 遍历模板
// 遍历是很多能力的基础
// 这里只按遍历顺序打印节点值，实际上可以做任何事情

// 前序：根左右
// 中序：左根右
// 后序：左右根

// 递归前序遍历
func PreOrder(root *TreeNode) {
	// 退出条件
	if root == nil {
		return
	}
	fmt.Println(root.Val) // 这里可以对根节点做任何事情，指当前子树的根节点，每个节点都有可能是子树的根节点
	PreOrder(root.Left)
	PreOrder(root.Right)
}

// 递归中序遍历
func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Println(root.Val)
	InOrder(root.Right)
}

// 递归后序遍历
func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Println(root.Val)
}

// 分层遍历
func LevelOrder(root *TreeNode) {
	if root != nil {
		return
	}
	queue := structure.NewQueue[*TreeNode]()
	queue.Push(root)
	for queue.Len() > 0 {
		length := queue.Len()
		// 遍历队列, 弹出节点, 每弹出一个把左右孩子插入队尾
		// 一次内部循环就是一层
		for i := 0; i < length; i++ {
			node := queue.Pop()
			fmt.Print(node.Val, " ")
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		fmt.Println()
	}
	return
}

// 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
func LevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		count := len(queue)
		var seq []int
		for i := 0; i < count; i++ {
			node := queue[i]
			seq = append(seq, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[count:]
		res = append(res, seq)
	}
	return res
}

// 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
// Z字型遍历
func LevelOrderZ(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		count := len(queue)
		var seq []int
		for i := 0; i < count; i++ {
			node := queue[i]
			seq = append(seq, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[count:]
		res = append(res, seq)
	}
	for i := range res {
		if i%2 == 1 {
			reverseSlice(res[i])
		}
	}
	return res
}

// 数组反转
func reverseSlice(slice []int) {
	i := 0
	j := len(slice) - 1
	for i < j {
		sort.Swap(slice, i, j)
		i++
		j--
	}
}
