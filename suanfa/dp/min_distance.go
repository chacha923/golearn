package dp

import "golearn/suanfa/lib"

// https://leetcode-cn.com/problems/edit-distance/
// 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
// 你可以对一个单词进行如下三种操作：
// 插入一个字符
// 删除一个字符
// 替换一个字符
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')

// base case 是 i 走完 s1 或 j 走完 s2，可以直接返回另一个字符串剩下的长度。
// 对于每对儿字符 s1[i] 和 s2[j]，可以有四种操作：
// if s1[i] == s2[j]:
//     啥都别做（skip）
//     i, j 同时向前移动
// else:
//     三选一：
//         插入（insert）
//         删除（delete）
//         替换（replace）
func minDistances(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 编辑空字符串
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] //不替换
			} else {
				dp[i][j] = dp[i-1][j-1] + 1 //替换
			}
			dp[i][j] = lib.Min(dp[i][j], dp[i][j-1]+1) // 插入
			dp[i][j] = lib.Min(dp[i][j], dp[i-1][j]+1) // 删除
		}
	}
	return dp[m][n]
}
