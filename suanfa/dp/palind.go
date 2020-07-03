package dp

import "golearn/suanfa/lib"

// 最长回文子串
// 动态规划方法, P(i,j) = true (子串Si-Sj是回文) | false 其他情况

// 中心展开法
// 事实上，只需使用恒定的空间，我们就可以在 O(n^2)O(n2) 的时间内解决这个问题。
// 我们观察到回文中心的两侧互为镜像。因此，回文可以从它的中心展开，并且只有 2n - 12n−1 个这样的中心。
// 你可能会问，为什么会是 2n - 12n−1 个，而不是 nn 个中心？原因在于所含字母数为偶数的回文的中心可以处于两字母之间（例如 \textrm{“abba”}“abba” 的中心在两个 \textrm{‘b’}‘b’ 之间）。

func longestPalindrome(s string) string {
	var start int
	var length int                //子串长度
	for i := 0; i < len(s); i++ { //奇数长度子串
		left := i - 1
		right := i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		if right-left-1 > length {
			start = left + 1
			length = right - left - 1
		}
	}
	for i := 0; i < len(s); i++ { //偶数长度
		left := i
		right := i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		if right-left-1 > length {
			start = left + 1
			length = right - left - 1
		}
	}
	return s[start : start+length]
}

// todo: 动态规划法

// 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
// 示例 1:
// 输入: "(()"
// 输出: 2
// 解释: 最长有效括号子串为 "()"
// 示例 2:
// 输入: ")()())"
// 输出: 4
// 解释: 最长有效括号子串为 "()()"
// 注意是子串!!!
// 用栈解决
// 对于遇到的每个 '(' ，我们将它的下标放入栈中。
// 对于遇到的每个 ')' ，我们弹出栈顶的元素并将当前元素的下标与弹出元素下标作差，得出当前有效括号字符串的长度。通过这种方法，我们继续计算有效子字符串的长度，并最终返回最长有效子字符串的长度。
func longestValidParentheses(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	stack := make([]int, 0) // 存下标, 不要存字符
	maxNum := 0             // 长度
	start := -1             //为语义一致即栈内只有'('的坐标，使用start标记起始位置。连续括号开始的位置前一个
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i) //'('入栈
		} else if s[i] == ')' {
			// // 弹出
			// stack = stack[:len(stack)-1]
			if len(stack) == 0 { // 遇到栈空了, 说明没有可配对的'('
				start = i
			} else {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					//空了 说明前面有可能有已经消除的配对括号，所以用start计算大小。
					maxNum = lib.Max(maxNum, i-start)
				} else {
					maxNum = lib.Max(maxNum, i-stack[len(stack)-1])
				}
			}
		}
	}
	return maxNum
}

// dp解决
// 我们定义一个 dp 数组，其中第 i 个元素表示以下标为 i 的字符结尾的最长有效子字符串的长度。
// 以 ( 结尾的子字符串不考虑，因为不可能构成合法括号
// i - 2 有可能小于零越界了，这种情况下就是只有 () ，前面记为 0 就好了.
// i - dp[i - 1] - 1 和 i - dp[i - 1] - 2 都可能越界，越界了当成 0 来计算就可以了.

// 如果 s[i] 为 ')', 那么 考虑两种情况：
// 第一种情况：如果 s[i - 1] 为 '(', 那么刚好可以和 s[i] 匹配， f[i + 1] = f[i - 1] + 2;
// 第二种情况：如果 s[i - 1] 为 ')', 那么只有 f[i]不为0（即以s[i - 1]结尾有效括号的长度），
// ###才 有可能使得以s[i]为结尾的有可能组成有效的括号，此时 f[i + 1] = f[i] + f[i - f[i] - 1] + 2。
// https://leetcode-cn.com/problems/longest-valid-parentheses/solution/dong-tai-gui-hua-si-lu-xiang-jie-c-by-zhanganan042/
func longestValidParenthesesWithDP(s string) int {
	maxNum := 0
	dp := make([]int, len(s)) // 默认0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// 处理边界
				if i < 2 {
					dp[i] = 2
				} else {
					dp[i] = dp[i-2] + 2 // 多了一对()
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' { // s[i-1] == ')'
				// ((...)) 此时一定是这种形式, i-1位置必须是有效括号对, 否则s[i]一定无法匹配
				// i-dp[i-1]-1  -> 与 s[i]匹配的位置
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
				} else { // 下标溢出了, 直接忽略
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxNum = lib.Max(dp[i], maxNum)
	}
	return maxNum
}

// 最大正方形面积
// 我们用 dp(i, j)dp(i,j) 表示以 (i,j) 为右下角，且只包含 1 的正方形的边长最大值
// 如果该位置的值是 11，则 dp(i, j)dp(i,j) 的值由其上方、左方和左上方的三个相邻位置的 dp 值决定。具体而言，当前位置的元素值等于三个相邻位置的元素中的最小值加 1，状态转移方程如下：
// 证明: https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/solution/tong-ji-quan-wei-1-de-zheng-fang-xing-zi-ju-zhen-2/
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSide := 0 // 记录最大边长值
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			// matrix[i][j] == 1, dp至少为1, 顺便处理边界
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = lib.Min(lib.Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}
	return maxSide * maxSide
}
