package dp

//求数组的最长递增子序列 (lis)  1.滑动窗口法 2.动态规划
//另一种解法: 拷贝数组并排序, 转化为求最长公共子序列lcs 的 长度
//状态转移方程:
// d(i) = max{d(j)+1, 1}
// 当 arr[i] > arr[j], 那么 dp[i] = dp[j] + 1, 否则 dp[i] = 1; 其中 j 是小于 i 的最大下标

// 用大白话解释就是，想要求d(i)，就把i前面的各个子序列中，
// 最后一个数不大于A[i]的序列长度加1，然后取出最大的长度即为d(i)。
// 当然了，有可能i前面的各个子序列中最后一个数都大于A[i]，那么d(i)=1， 即它自身成为一个长度为1的子序列。

var res = make([]int, 0)

//求 数组的 lis 长度
func lisLength(arr []int, length int) int {
	// dp[i] 表示以第 i 个元素为结尾的最长递增子序列的长度
	dp := make([]int, length)
	//init 初始化, 子数组的lis值默认为1
	for k := range dp {
		dp[k] = 1
	}

	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = dp[j] + 1
			}
		}
	}

	//找出dp[]的最大值
	max := dp[0]
	for i := 1; i < length; i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}

	return max
}

func lisLengthWindow(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}
	// count 为当前元素峰值，ans为最大峰值
	count := 1
	ans := 1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] > arr[i] {
			count++
		} else {
			count = 1
		}
		if count > ans {
			ans = count
		}
	}
	return ans
}
