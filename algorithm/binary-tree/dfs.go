package binary_tree

import (
	"fmt"
	"golearn/algorithm/structure"
	"golearn/algorithm/util"
)

// 非递归的前中后序遍历

//1.首先将根节点放入队列中。
//2.从队列中取出第一个节点，并检验它是否为目标。
//  如果找到目标，则结束搜寻并回传结果。
//  否则将它某一个尚未检验过的直接子节点加入队列中。
//3.重复步骤2。
//4.如果不存在未检测过的直接子节点。
//  将上一级节点加入队列中。
//  重复步骤2。
//5.重复步骤4。
//6.若队列为空，表示整张图都检查过了——亦即图中没有欲搜寻的目标。结束搜寻并回传“找不到目标”。

// 需要队列或栈辅助 (回溯思想)
// 前序遍历
func PreOrder1(root *TreeNode) {
	if root == nil {
		return
	}
	stack := structure.NewStack[*TreeNode]()
	curr := root
	stack.Push(curr) // 根节点进栈
	for stack.Len() != 0 {
		// 出栈, 开始操作
		curr = stack.Pop()
		fmt.Println(curr.Val)
		// 进栈, 注意栈的特性, 先进右后进左
		if curr.Right != nil {
			stack.Push(curr.Right)
		}
		if curr.Left != nil {
			stack.Push(curr.Left)
		}
	}
	return
}

// 中序遍历
func InOrder1(root *TreeNode) {
	if root == nil {
		return
	}
	stack := structure.NewStack[*TreeNode]()
	curr := root
	for curr != nil || stack.Len() != 0 {
		// 进栈
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}
		// 出栈
		curr = stack.Pop()
		curr = curr.Right
	}
	return
}

// 后序遍历
// 骚技巧, 把前序遍历的逆序反过来
// 前序遍历根->左->右， 后序遍历 左->右->根
// 将前序遍历改成 根->右->左
// 反转即可得到后序遍历 (这只能打印)
func PostOrder1(root *TreeNode) {
	if root == nil {
		return
	}
	stack := structure.NewStack[*TreeNode]()
	m := make(map[*TreeNode]struct{}) // 记录已经访问的结点
	curr := root
	stack.Push(curr) // 根节点进栈
	for stack.Len() > 0 {
		// 取栈顶
		curr = stack.Top()
		leftVisited, rightVisited := true, true // 标记左右孩子是否被访问
		// 进栈
		if curr.Right != nil {
			if _, ok := m[curr.Right]; !ok {
				rightVisited = false
				// 右孩子没有访问过，进栈
				stack.Push(curr.Right)
			}
		}
		if curr.Left != nil {
			if _, ok := m[curr.Left]; !ok {
				leftVisited = false
				// 左孩子没有访问过，进栈
				stack.Push(curr.Left)
			}
		}
		// 遇到叶子节点, 出栈
		if leftVisited && rightVisited {
			m[curr] = struct{}{}
			top := stack.Pop()
			fmt.Println(top.Val)
		}
	}
	return
}

// 递归模版 深度优先遍历（前序）
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	// Do something
	fmt.Println(root.Val)
	dfs(root.Left)
	dfs(root.Right)
}

// 左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		ans += root.Left.Val
	}
	if root.Left != nil {
		ans += sumOfLeftLeaves(root.Left)
	}
	if root.Right != nil {
		ans += sumOfLeftLeaves(root.Right)
	}
	return ans
}

// 给定一棵二叉树，以逆时针顺序从根开始返回其边界。边界按顺序包括左边界、叶子结点和右边界而不包括重复的结点。 (结点的值可能重复)
// dfs + 左右边界标记
// 简单理解: 一个先序遍历, 给节点标记flag
// https://leetcode-cn.com/problems/boundary-of-binary-tree/solution/er-cha-shu-de-bian-jie-by-leetcode/
func boundaryOfBinaryTree(root *TreeNode) []int {
	var res = make([]int, 0)
	boundary(root, true, true, &res)
	return res
}

func boundary(node *TreeNode, leftBound, rightBound bool, res *[]int) {
	if node == nil {
		return
	}
	if leftBound {
		*res = append(*res, node.Val)
	} else if node.Left == nil && node.Right == nil {
		*res = append(*res, node.Val)
		return
	}
	boundary(node.Left, leftBound, !leftBound && rightBound && node.Right == nil, res)
	boundary(node.Right, !rightBound && leftBound && node.Left == nil, rightBound, res)
	if !leftBound && rightBound {
		*res = append(*res, node.Val)
	}
}

// 求最大深度
func MaxDepth(head *TreeNode) int {
	if head == nil {
		return 0
	}
	left := MaxDepth(head.Left)
	right := MaxDepth(head.Right)
	//注意退出条件，取左右子树的最大深度，再加上 1（本节点）
	return util.Max(left, right) + 1
}

var maxDepth int
var curDepth int

// 用回溯思想求最大深度
func MaxDepthBackstrack(head *TreeNode) int {
	var traverse = func(head *TreeNode) {
		if head == nil {
			return
		}
		curDepth++
		if head.Left == nil && head.Right == nil {
			maxDepth = util.Max(maxDepth, curDepth)
		}
		MaxDepthBackstrack(head.Left)
		MaxDepthBackstrack(head.Right)
		curDepth--
	}

	traverse(head)
	return maxDepth
}

// 最小深度
func MinDepth(head *TreeNode) int {
	if head == nil {
		return 0
	}
	// null节点不参与比较
	if head.Left != nil && head.Right == nil {
		return 1 + MinDepth(head.Left)
	}
	if head.Left == nil && head.Right != nil {
		return 1 + MinDepth(head.Right)
	}
	return util.Min(MinDepth(head.Left), MinDepth(head.Right)) + 1
}
