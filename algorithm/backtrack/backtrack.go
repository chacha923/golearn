package backtrack

import (
	binary_tree "golearn/algorithm/binary-tree"
	"strconv"
	"strings"
)

// 回溯算法也是 dfs 思想的应用
//「回溯」就是 深度优先遍历 状态空间的过程中发现的特有的现象，程序会回到以前访问过的结点。
// 而程序在回到以前访问过的结点的时候，就需要将状态变量恢复成为第一次来到该结点的值。（跳出递归，后退一步）

// 在代码层面上，在递归方法结束以后，执行递归方法之前的操作的 逆向操作 即可。

// 括号生成, 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := make([]string, 0)
	backtrackParenthesis(&res, "", 0, 0, n)
	return res
}

// curStr: 当前字符串
// open: curStr 左括号数量
// close: curStr 右括号数量
func backtrackParenthesis(res *[]string, curStr string, open int, close int, n int) {
	if len(curStr) == n*2 {
		// 符合满足条件的 n 对括号，添加到结果集
		*res = append(*res, curStr)
		return
	}

	if open < n {
		// 左括号数量小于 n，可以添加左括号
		backtrackParenthesis(res, curStr+"(", open+1, close, n)
		// 后退
		curStr = strings.TrimSuffix(curStr, "(")
	}
	if close < open {
		// 右括号数量小于左括号数量，可以添加右括号
		backtrackParenthesis(res, curStr+")", open, close+1, n)
		// 后退
		curStr = strings.TrimSuffix(curStr, ")")
	}
}

// tmp 当前字符串
// left ( 待匹配个数
// right ) 待匹配个数
// 回溯算法，回溯跳出条件就是左右括号都已经排完的情况。
// 括号成对存在，先有左括号再有右括号，所以只有右括号的数量小于左括号才进行右括号的添加。
// 最后如果右括号的数量等于0，表示右括号已经排完了，同时意味着左括号也排完了。
func backtrack(left, right int, tmp string, res *[]string) {
	/*
	   回溯跳出条件，
	   并不需要判断左括号是否用完，因为右括号生成的条件 right > left ，
	   所以右括号用完了就意味着左括号必定用完了
	*/
	if right == 0 {
		*res = append(*res, tmp)
	}

	// 生成左括号
	if left > 0 {
		backtrack(left-1, right, tmp+"(", res)
	}

	// 括号成对存在，有左括号才会有右括号
	if right > left {
		backtrack(left, right-1, tmp+")", res)
	}
}

// 复原ip地址
// 转化为放3个点的问题, 由于返回全部的解, 考虑回溯剪枝
func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}
	result := new([]string)
	back(s, 0, []string{}, result)
	return *result
}

func back(s string, pos int, cur []string, ans *[]string) {
	// 找出4段了, 此时pos也刚好遍历完s, 返回
	if len(cur) == 4 {
		if pos == len(s) {
			*ans = append(*ans, strings.Join(cur, "."))
		}
		return
	}
	// 每个ip段最多有3个数字
	for i := 1; i <= 3; i++ {
		// 如果当前位置距离 s 末尾小于 3 就不用再分段了，直接跳出循环即可。
		if pos+i > len(s) {
			break
		}
		segment := s[pos : pos+i]
		// 剪枝条件: 起始为0且长度>1 , 不能大于255
		if strings.HasPrefix(segment, "0") && len(segment) > 1 {
			continue
		}
		segmentNum, _ := strconv.Atoi(segment)
		if segmentNum > 255 {
			continue
		}

		// 符合要求就加入到 cur 数组中
		cur = append(cur, segment)
		// 继续递归遍历下一个位置
		back(s, pos+i, cur, ans)
		// 回退到上一个元素，即回溯
		cur = cur[:len(cur)-1]
	}
}

// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 解集不能包含重复的「子集」。
// [1,2,3] -> [],[1],[2],[3],[1,2],[1,3],[2,3],[1,2,3]
func subsets(nums []int) [][]int {
	var track = make([]int, 0) // 记录当前走过的路径
	var res = make([][]int, 0) // 结果集
	// back1(nums, track, 0, res)

	// 定义回溯函数，从指定下标开始的子数组，统计子集
	var backtrack func(start int)
	backtrack = func(start int) {
		// 每次加入一个子集
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}

// 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
func pathSum(root *binary_tree.TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	path := make([]int, 0) // 储存路径value
	res := new([][]int)    // 结果集
	num := 0               // 和
	back2(root, num, sum, path, res)
	return *res
}

func back2(root *binary_tree.TreeNode, num int, sum int, path []int, res *[][]int) {
	// 退出条件, num == sum, 访问到叶子节点
	if root == nil {
		if num == sum {
			tmp := []int{}
			tmp = append(tmp, path...)
			*res = append(*res, tmp)
		}
		return
	}
	path = append(path, root.Val)
	if root.Left == nil {
		// 左子树空, 剪枝
		back2(root.Right, num+root.Val, sum, path, res)
	} else if root.Right == nil {
		// 右子树空, 剪枝
		back2(root.Left, num+root.Val, sum, path, res)
	} else {
		back2(root.Left, num+root.Val, sum, path, res)
		back2(root.Right, num+root.Val, sum, path, res)
	}
	// 撤销
	num -= path[len(path)-1]
	path = path[:len(path)-1]
}

// 判断是否存在等于目标和的路径
func hasPathSumBackstrack(root *binary_tree.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var curSum int
	var found bool

	var back func(root *binary_tree.TreeNode)
	back = func(root *binary_tree.TreeNode) {
		if root == nil {
			return
		}
		// 前序遍历位置
		curSum += root.Val
		// 到叶子节点了，和 target 比较
		if root.Left == nil && root.Right == nil {
			if curSum == targetSum {
				found = true
			}
		}
		back(root.Left)
		back(root.Right)
		// 后续遍历位置
		curSum -= root.Val
	}
	back(root)
	return found
}

// 简化版, 递归解法，只判断是否存在等于目标和的路径
// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
func hasPathSum(root *binary_tree.TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	// 叶子节点
	if root.Left == nil && root.Right == nil {
		if sum-root.Val == 0 {
			return true
		} else {
			return false
		}
	}
	// sum - 当前节点值，等于左子树或右子树的路径和，则找到
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func SubSets(nums []int) [][]int {
	res := [][]int{} // 结果集
	track := []int{} // 临时数组，可以理解成栈

	// start 表示从下标 start 开始到末尾，nums 片段的排列
	var backtrack func(start int)
	backtrack = func(start int) {
		// 拷贝 track，放到结果集
		var temp = make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			// 做选择
			track = append(track, nums[i])
			// 递归，进入下一个状态
			backtrack(i + 1)
			// 撤销选择，track 回退
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return res
}
