package dp

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
