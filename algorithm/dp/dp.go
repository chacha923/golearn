package dp

import (
	"golearn/algorithm/util"
	"math"
)

// 爬楼梯，一次可以爬1级或2级，问爬到第n级有多少种方法
func stair(n int) int {
	// dp[n] = dp[n-1] + dp[n-2]
	// 最小问题 不能再分解
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return stair(n-1) + stair(n-2)
}

// not recursive
func stair2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 给定一个包含非负整数的 m x n 网格 grid
// 请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// 1 3 1
// 1 5 1
// 4 2 1
// => 1 3 1 1 1 ==> 7
func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var width = len(grid[0])
	var height = len(grid)
	var dp = make([][]int, height)

	//	初始化 dp
	for i := 0; i < height; i++ {
		dp[i] = make([]int, width)
	}
	for i := 1; i < height; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < width; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	// 状态转移方程
	// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
	for i := 1; i < height; i++ {
		for j := 1; j < width; j++ {
			dp[i][j] = util.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[height-1][width-1]
}

// 0 1 背包
// 有一个背包，最大容量为w，现有n种不同的物品，编号为0...n-1，其中每一件物品的重量为w(i)，价值为v(i)
func knapspack(w int, n int, wt []int, val []int) int {
	// 状态：背包剩余的空间有多少，物品还剩多少
	// 选择：装进背包，不装进背包

	// dp 数组定义：dp[i][j] 前 i 个物品，背包已占用容量为 j 时，可以装的最大价值
	// 举例：dp[3][5] = 6, 前 3 个物品，背包已占用容量为 5 时，可以装的最大价值为 6
	// base case dp[0][..] = dp[..][0] = 0

	var dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			// 容量约束，背包装不下，只能不装
			if j-wt[i] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				// 能装，考虑装或者不装，取最大值
				dp[i][j] = util.Max(
					dp[i-1][j-wt[i]]+val[i], // 装，就用装之前的最大值加上当前物品的价值
					dp[i-1][j],              //不装
				)
			}
		}
	}
	return dp[n][w]
}

// 将 s1 转换成 s2 的最小操作数
// 插入 删除 替换
// eg: s1 = "horse", s2 = "ros"，最小操作数为 3：将 h 替换成 r，删除 r，删除 e
// 也可以用自顶向下递归+备忘录的方式
func minDistance(s1, s2 string) int {
	// 这里 s1[i] 表示 s1 的第 i 个字符，s2[j] 表示 s2 的第 j 个字符
	// dp[i][j] 表示 s1[0...i] 转换成 s2[0...j] 的最小操作数
	var dp = make([][]int, len(s1)+1)
	for i := 1; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2)+1)
		dp[i][0] = i // base case
	}
	for j := 1; j <= len(s2); j++ {
		dp[0][j] = j // base case
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			// 四种选择
			if s1[i-1] == s2[j-1] {
				// s1[i] 和 s2[j] 相等，不需要操作
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i][j-1]+1,   // 插入
					dp[i-1][j]+1,   // 删除
					dp[i-1][j-1]+1, // 替换
				)
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func min(a, b, c int) int {
	var min = util.Min(a, b)
	return util.Min(min, c)
}

// 给你输入一个整数数组 nums，请你找在其中找一个和最大的子数组，返回这个子数组的和
// 1. 滑动窗口法 2. 动态规划
// 考虑 nums 包含负数的情况
func maxSubArrayByWindow(nums []int) int {
	var (
		left      = 0
		right     = 0 // 窗口边界
		windowSum = 0 // 窗口内元素的和
		maxSum    = math.MinInt
	)
	// 我们可以在窗口内元素之和大于等于 0 时扩大窗口，在窗口内元素之和小于 0 时缩小窗口，在每次移动窗口时更新答案。
	// labuladong 起初认为 nums 包含负数不能使用滑动窗口，但其实可以
	for right < len(nums) {
		windowSum += nums[right]
		right++
		if windowSum > maxSum {
			maxSum = windowSum
		}
		// 窗口内元素和小于0，判断是否要收缩
		for windowSum < 0 {
			windowSum -= nums[left]
			left++
		}
	}
	return maxSum
}

func maxSubArray(nums []int) int {
	// 怎么由 dp[n-1] 推导出 dp[n] 呢？
	// 定义 dp[i] 为以 nums[i] 结尾的最大子序和，因此每次要遍历整个 dp 数组，找出最大值

	var res = math.MinInt
	var dp = make([]int, len(nums))

	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// 要么自成一派，要么和前面的子数组合并
		dp[i] = util.Max(dp[i-1]+nums[i], nums[i])
	}
	// 找 dp 数组中的最大值
	for i := 0; i < len(dp); i++ {
		res = util.Max(res, dp[i])
	}
	return res
}

// 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断达到最后一个下标的最小跳跃次数
// 只能穷举了
func jump(nums []int) int {
	var memo = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		// 最大跳跃次数 肯定不会超过 len(nums)
		memo[i] = len(nums)
	}

	// 从索引 p 跳到最后一格，至少需要 dp(nums, p) 步
	var dp func(memo []int, nums []int, p int) int
	dp = func(memo []int, nums []int, p int) int {
		n := len(nums)
		// base case: 就是当 p 超过最后一格时，不需要跳跃：
		if p >= n-1 {
			return 0
		}
		// 子问题已经计算过
		if memo[p] != n {
			return memo[p]
		}
		// 穷举每个选择
		var steps = nums[p]
		// 你可以从 p 位置选择跳 1 步，2 步...
		for i := 1; i <= steps && p <= n-1; i++ {
			sub := dp(memo, nums, p+i)
			memo[p] = util.Min(memo[p], sub+1)
		}
		return memo[p]
	}

	dp(memo, nums, 0)
}
