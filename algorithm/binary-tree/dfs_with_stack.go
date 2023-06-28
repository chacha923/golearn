package binary_tree

import (
	"fmt"
	"golearn/algorithm/structure"
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
		// do something
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
		// TODO
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
			// do something
			fmt.Println(top.Val)
		}
	}
	return
}
