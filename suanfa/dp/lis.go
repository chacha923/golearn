package dp

import "fmt"

//求数组的﻿最长递增子序列 (lis)

//另一种解法: 拷贝数组并排序, 转化为求最长公共子序列lcs 的 长度

//状态转移方程:
// d(i) = max{d(j)+1, 1} (当且仅当  j < i , arr[j] < arr[i])

// ﻿用大白话解释就是，想要求d(i)，就把i前面的各个子序列中，
// 最后一个数不大于A[i]的序列长度加1，然后取出最大的长度即为d(i)。
// 当然了，有可能i前面的各个子序列中最后一个数都大于A[i]，那么d(i)=1， 即它自身成为一个长度为1的子序列。

var res = make([]int, 0)

//求 数组的 lis 长度
func lisLength(arr []int, length int) {
	lis := make([]int, 0, length) //lis数组保存 0~index 子数组的 lis长度
	//init 初始化, 子数组的lis值默认为1
	for k := range lis {
		lis[k] = 1
	}

	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}

	//找出lis[]的最大值
	max := lis[0]
	for i := 1; i < length; i++ {
		if max < lis[i] {
			max = lis[i]
		}
	}

	fmt.Println(max)
}
