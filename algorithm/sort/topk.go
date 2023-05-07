package sort

// 4种解法秒杀TopK（快排/堆/二叉搜索树/计数排序）
// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/solution/3chong-jie-fa-miao-sha-topkkuai-pai-dui-er-cha-sou/

func getLeastNumbersWithCountSort(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	tmp := make([]int, 10001)
	for i := range arr {
		tmp[arr[i]]++ //统计每个数出现的次数
	}
	result := make([]int, 0)
	for i := 0; i < len(tmp); i++ {
		c := tmp[i]
		if c == 0 {
			continue
		}
		for j := 0; j < c; j++ {
			result = append(result, i)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}

func getLeastNumbersWithFastSort(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	return nil
}

// 快速排序 提前返回
func fastSort(arr []int, start int, end int, k int) []int {
	if start < end {
		i := start
		j := end
		mid := arr[i] //取第一个元素为mid元素
		for i < j {
			for arr[j] > mid && i < j { // 从右向左找第一个小于等于mid的数
				j--
			}
			for arr[i] < mid && i < j { // 从左向右找第一个大于等于mid的数
				i++
			}
			if i <= j {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
		//arr[i] = mid	//此时arr[i] 已经位于最终位置了
		if i == len(arr)-k {
			return arr[i:]
		}
		if i > len(arr)-k {
			FastSort(arr, start, i-1)
		}
		if i < len(arr)-k {
			FastSort(arr, j+1, end)
		}
	}
	return arr
}
