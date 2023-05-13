package dp

import "golearn/algorithm/util"

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
