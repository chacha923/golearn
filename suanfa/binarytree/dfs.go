package binarytree

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
	stack := make([]*TreeNode, 0)
	curr := root
	stack = append(stack, curr) // 根节点进栈
	for len(stack) != 0 {
		// 出栈, 开始操作
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 进栈, 注意栈的特性, 先进右后进左
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return
}

// 中序遍历
func InOrder1(root *TreeNode) {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	curr := root
	for curr != nil || len(stack) != 0 {
		// 进栈
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		// 出栈
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
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
	stack := make([]*TreeNode, 0)
	m := make(map[*TreeNode]struct{}) // 记录已经访问的结点
	curr := root
	stack = append(stack, curr) // 根节点进栈
	for len(stack) > 0 {
		// 取栈顶
		curr = stack[len(stack)-1]
		leftVisited, rightVisited := true, true // 标记左右孩子是否被访问
		// 进栈
		if curr.Right != nil {
			if _, ok := m[curr.Right]; !ok {
				rightVisited = false
				stack = append(stack, curr.Right)
			}
		}
		if curr.Left != nil {
			if _, ok := m[curr.Left]; !ok {
				leftVisited = false
				stack = append(stack, curr.Left)
			}
		}
		// 遇到叶子节点, 出栈
		if leftVisited && rightVisited {
			m[curr] = struct{}{}
			stack = stack[:len(stack)-1]
		}
	}
	return
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
	res = new([]int)
	dfs1(root, true, true, res)
	return res
}

func dfs1(node *TreeNode, leftBound, rightBound bool, res *[]int) {
	if root == nil {
		return
	}
	if leftBound {
		*res = append(*res, node.Val)
	} else if node.Left == nil && node.Right == nil {
		*res = append(*res, node.Val)
		return
	}
	dfs1(node.Left, leftBound, !leftBound && rightBound && node.Right == nil, res)
	dfs1(node.Right, !rightBound && leftBound && node.Left == nil, rightBound, res)
	if !leftBound && rightBound {
		*res = append(*res, node.Val)
	}
}

// class Solution {
//     public List<Integer> boundaryOfBinaryTree(TreeNode root) {
//         List<Integer> res = new ArrayList<>();
//         dfs(root, true, true, res);
//         return res;
//     }

//     private void dfs(TreeNode node, boolean leftBound, boolean rightBound, List<Integer> res) {
//         if (node == null) {
//             return;
//         }
//         if (leftBound) {
//             res.add(node.val);
//         } else if (node.left == null && node.right == null) {
//             res.add(node.val);
//             return;
//         }
//         dfs(node.left, leftBound, !leftBound && rightBound && node.right == null, res);
//         dfs(node.right, !rightBound && leftBound && node.left == null, rightBound, res);
//         if (!leftBound && rightBound) {
//             res.add(node.val);
//         }
//     }
// }
