package dp

import "golearn/suanfa/lib"

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

// dp[i][j]
// 第一维 i 表示索引为 i 的那一天（具有前缀性质，即考虑了之前天数的收益）能获得的最大利润；
// 第二维 j 表示索引为 i 的那一天是持有股票，还是持有现金。这里 0 表示持有现金（cash），1 表示持有股票（stock）。
func maxProfit(prices []int) int {
	// 特殊情况, 0天或1天
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2) // 只记录持有现金或股票
	}
	dp[0][0] = 0          // 如果什么都不做
	dp[0][1] = -prices[0] // 如果买入股票，当前收益是负数

	for i := 1; i < len(prices); i++ {
		dp[i][0] = lib.Max(dp[i-1][0], dp[i-1][1]+prices[i]) // 第i天持币, 要么第i-1天持币, 或者第i-1天持股第i天卖出
		dp[i][1] = lib.Max(dp[i-1][1], dp[i-1][0]-prices[i]) // 第i天持股, 要么第i-1天持股, 或者第i-1天持币第i天买入
	}

	return dp[len(prices)-1][0] //最后一天必须持有现金
}

// 最简单的一次遍历, 转化为等价问题, 只求结果不能得过程
func maxProfitSimple(prices []int) int {

}
