package binary_tree

import (
	"fmt"
	"golearn/algorithm/structure"
)

// 不用递归的方式实现遍历
// 实际是考察对递归的理解，自己实现栈（迭代法）
// 递归的本质是栈，所以可以用栈来实现

// 先序遍历
func PreOrderWithoutRecurve(root *TreeNode) {
	if root == nil {
		return
	}
	stack := structure.NewStack[*TreeNode]()
	stack.Push(root)
	for stack.Len() != 0 {
		var node = stack.Pop()
		// do something
		fmt.Println(node.Val)
		// 这里要注意，先压右子树，再压左子树
		// 因为栈的特性是先进后出，所以先压右子树，后压左子树，这样出栈的时候就是先左后右
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}

// 中序遍历是左中右，先访问的是二叉树顶部的节点，然后一层一层向下访问，直到到达树左面的最底部，再开始处理节点
// 这就造成了处理顺序和访问顺序是不一致的。
// 那么在使用迭代法写中序遍历，就需要借用指针的遍历来帮助访问节点，栈则用来处理节点上的元素。
func InOrderWithoutRecurve(root *TreeNode) {
	if root == nil {
		return
	}
	stack := structure.NewStack[*TreeNode]()
	var current *TreeNode = root // 保存当前节点
	for current != nil || stack.Len() > 0 {
		if current != nil {
			// 指针来访问节点，访问到最底层
			// 不断把根和左子树压入栈
			stack.Push(root)
			root = root.Left
			continue
		}
		current = stack.Pop()
		// do something
		fmt.Println(current.Val)
		// 弹出后，下一轮把右子树压入栈
		current = current.Right
	}
}

// 后序遍历
func PostOrderWithoutRecurve(root *TreeNode) {
	if root == nil {
		return
	}
	var stack = structure.NewStack[*TreeNode]()
	var prev *TreeNode // 保存上一个访问的节点，用于判断右子树是否已经访问过

	for root != nil || stack.Len() > 0 {
		for root != nil {
			// 根和左节点不断压栈
			stack.Push(root)
			root = root.Left
		}
		var top = stack.Top()
		// 如果栈顶节点的右子树存在且未被访问过，则不出栈，继续下一轮循环使右子树入栈
		if top.Right != nil && top.Right != prev {
			root = top.Right
		} else {
			// do something
			fmt.Println(top.Val)
			prev = top
		}
	}
}
