package other

import "golearn/algorithm/util"

// 贪心

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
func maxArea(height []int) int {
	var left = 0
	var right = 0
	var maxArea = 0

	for left < right {
		var area = (right - left) * util.Min(height[left], height[right])
		maxArea = util.Max(maxArea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标。
func jump(nums []int) bool {
	var (
		n        = len(nums)
		farthest = 0
	)
	for i := 0; i < n; i++ {
		// 不断计算能跳到的最远距离
		farthest = util.Max(farthest, i+nums[i])
		// 无法到达下标为 i 的位置，直接返回
		// 可能碰到了 0，卡住跳不动了
		if farthest <= i {
			return false
		}
	}
	return farthest >= n-1
}
