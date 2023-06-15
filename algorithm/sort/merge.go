package sort

import "fmt"

func RunMergeSort() {
	tmp := make([]int, len(arr)) //在排序前，先建好一个长度等于原数组长度的临时数组，避免递归中频繁开辟空间
	mergeSort(arr, tmp, 0, len(arr)-1)
	fmt.Println(arr)
}

/**
 * 归并排序
 * 简介:将两个（或两个以上）有序表合并成一个新的有序表 即把待排序序列分为若干个子序列，
 *     每个子序列是有序的。然后再把有序子序列合并为整体有序序列
 * 时间复杂度为O(nlogn)
 * 稳定排序方式
 * @param nums 待排序数组
 * @return 输出有序数组
 */
func sort(nums []int, lo int, hi int) {
	mid := lo + (hi-lo)/2
	sort(nums, lo, mid)
	sort(nums, mid+1, hi)

	merge(nums, lo, mid, hi)
}

func merge(nums []int, lo, mid, hi int) {
	// 对 nums[lo...mid] 和 nums[mid+1...hi] 合并
	// 创建临时数组，用于存储合并结果
	temp := make([]int, hi-lo+1)
	var i = lo      // 左侧子数组的起始索引
	var j = mid + 1 // 右侧子数组的起始索引
	var idx = 0     // 临时数组的索引

	// 将左右两个有序子数组按顺序合并到临时数组中
	for i <= mid && j <= hi {
		if nums[i] <= nums[j] {
			temp[idx] = nums[i]
			i++
		} else if nums[i] > nums[j] {
			temp[idx] = nums[j]
			j++
		}
		idx++
	}
	for i <= mid {
		temp[idx] = nums[i]
		i++
		idx++
	}
	for j <= hi {
		temp[idx] = nums[j]
		j++
		idx++
	}
	// 将临时数组的元素复制回原始数组
	for i := 0; i < len(temp); i++ {
		nums[lo+i] = temp[i]
	}
}

func RunMergeSort2() {
	arr = MergeSort2(arr)
	fmt.Println(arr)
}

func MergeSort2(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	key := n / 2
	left := MergeSort2(arr[0:key])
	right := MergeSort2(arr[key:])
	return merge2(left, right)
}

func merge2(left, right []int) []int {
	tmp := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	return tmp
}
