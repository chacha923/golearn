package backtrack

import (
	"strconv"
	"strings"
)

// 括号生成, 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
func generateParenthesis(n int) {
	res := new([]string)
	backtrack(n, n, "", res)
	return result
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
// 解集不能包含重复的子集。
func subsets(nums []int) [][]int {
	track := make([]int, 0) // 记录当前走过的路径
	res := new([][]int)     // 结果集
	back1(nums, 0, res)
}

func back1(nums []int, track []int, start int, res *[][]int) {
	*res = append(*res, track)
	// 注意 i 从 start 开始递增
	for i := start; i < len(nums); i++ {
		//做选择
		track = append(track, nums[i])
		// 回溯
		back1(nums, i+1, track, res)
		// 撤销
		track = track[:len(track)-1]
	}
	return
}

// 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	path := make([]int, 0) // 储存路径value
	res := new([][]int)    // 结果集
	num := 0               // 和
	back2(root, num, sum, path, res)
	return *res
}

func back2(root *TreeNode, num int, sum int, path []int, res *[][]int) {
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

// 简化版, 只判断是否存在等于目标和的路径
// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
func hasPathSum(root *TreeNode, sum int) bool {
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
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
