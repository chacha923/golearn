package other

import (
	"golearn/algorithm/structure"
	"golearn/algorithm/util"
)

// 接雨水

// 1. 动态规划
// 2. 双指针
// 3. 单调栈

// 对于每一个柱子接的水，那么它能接的水=min(左右两边最高柱子）-当前柱子高度，这个公式没有问题。同样的，两根柱子要一起求接水，同样要知道它们左右两边最大值的较小值。
func trap1(height []int) int {
	var left, right = 0, len(height) - 1
	// 已经遍历过的左右两侧最高柱子
	var leftMax, rightMax = 0, 0
	var ans = 0

	for left < right {
		leftMax = util.Max(leftMax, height[left])
		rightMax = util.Max(rightMax, height[right])

		if height[left] < height[right] {
			ans += leftMax - height[left]
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

// 普通栈
func trap2(height []int) int {
	var stack = structure.NewStack[int]()
	var sum = 0
	var current = 0

	for current < len(height) {
		// 如果栈不空并且当前指向的高度大于栈顶高度就一直循环
		for !stack.Empty() && height[current] > height[stack.Top()] {
			var idxOfTop = stack.Pop() // 取出要出栈的元素
			var heightOfTop = height[idxOfTop]
			if stack.Empty() {
				break // 栈空就出去
			}

			var distance = current - stack.Top() - 1 //两堵墙之前的距离。
			var min = util.Min(height[stack.Top()], height[current])
			sum += distance * (min - heightOfTop)
		}
		stack.Push(current) //当前指向的墙入栈
		current++           //指针后移
	}
	return sum
}

// 两次 dp，从左往右，从右往左
func trap3(height []int) int {
	var n = len(height)
	if n == 0 {
		return 0
	}
	// 记录下标为 i 时的最高点
	// 当 1≤i≤n−1 时，leftMax[i]=max⁡(leftMax[i−1],height[i])；
	// 当 0≤i≤n−2 时，rightMax[i]=max⁡(rightMax[i+1],height[i])。
	var leftMax = make([]int, n)
	var rightMax = make([]int, n)
	var ans = 0

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = util.Max(height[i], leftMax[i-1])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = util.Max(height[i], rightMax[i+1])
	}

	for i := 1; i < n-1; i++ {
		ans += util.Min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}
