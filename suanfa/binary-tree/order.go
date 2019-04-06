package binary_tree

import (
	"fmt"
	"golearn/suanfa/sort"

	"github.com/eapache/queue"
)

//遍历

//递归前序遍历
func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

//递归中序遍历
func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PreOrder(root.Left)
	fmt.Println(root.Val)
	PreOrder(root.Right)
}

//递归后序遍历
func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PreOrder(root.Left)
	PreOrder(root.Right)
	fmt.Println(root.Val)
}

//分层遍历
func LevelOrder(root *TreeNode) {
	if root != nil {
		return
	}
	queue := queue.New()
	queue.Add(root)
	for queue.Length() != 0 {
		length := queue.Length()
		for i := 0; i < length; i++ {
			node := queue.Remove().(*TreeNode)
			fmt.Print(node.Val, " ")
			if node.Left != nil {
				queue.Add(node.Left)
			}
			if node.Right != nil {
				queue.Add(node.Right)
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

func reverseSlice(slice []int) {
	i := 0
	j := len(slice) - 1
	for i < j {
		sort.Swap(slice, i, j)
		i++
		j--
	}
}
