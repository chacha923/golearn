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
