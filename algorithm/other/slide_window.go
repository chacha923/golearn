package other

import (
	"golearn/algorithm/structure"
	"golearn/algorithm/util"
)

// 最小覆盖子串
func minWindow(s string, t string) string {
	return ""
}

// 判断 s 中是否存在 t 的排列
func checkInclusion(t string, s string) bool {
	need := make(map[byte]int)   // t 中字符出现的次数
	window := make(map[byte]int) // 窗口中字符出现的次数
	for _, c := range []byte(t) {
		need[c]++
	}

	var left, right, valid int

	for right < len(s) {
		var c = s[right] // 当前新加入窗口的字符
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 循环条件，窗口长度大于等于 t 的长度
		for right-left >= len(t) {
			// 判断左边是否要收缩
			if valid == len(need) {
				// 找到了
				return true
			}
			var d = s[left] // 即将移出窗口的字符
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

// 寻找最长不重复子串
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	// 记录结果
	res := 0
	for right < len(s) {
		c := s[right]
		right++
		// 进行窗口内数据的一系列更新
		window[c]++
		// 判断左侧窗口是否要收缩
		for window[c] > 1 {
			d := s[left]
			left++
			// 进行窗口内数据的一系列更新
			window[d]--
		}
		// 在这里更新答案
		res = util.Max(res, right-left)
	}
	return res
}

// 给你输入一个数组 nums 和一个正整数 k，有一个大小为 k 的窗口在 nums 上从左至右滑动，请你输出每次窗口中 k 个元素的最大值。
func maxSlidingWindow(nums []int, k int) []int {
	window := structure.NewMonotonicQueue()
	var res []int

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			// 先把窗口的前 k-1 填满
			window.Push(nums[i])
		} else {
			// 窗口开始向前滑动
			window.Push(nums[i])
			// 将当前窗口中的最大元素记入结果
			res = append(res, window.Max())
			// 移出最后的元素
			window.Pop(nums[i-k+1])
		}
	}
	return res
}
