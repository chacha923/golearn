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
