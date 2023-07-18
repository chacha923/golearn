package sort

import "fmt"

func RunFastSort() {
	fmt.Println(arr)
	fmt.Println(FastSort(arr, 0, len(arr)-1))
}

// 快速排序，递归
func FastSort(arr []int, start int, end int) []int {
	if start < end {
		i := start
		j := end
		mid := arr[i] //取第一个元素为参照位
		for i < j {
			for arr[j] > mid && i < j { // 从右向左找第一个小于等于参照位的数
				j--
			}
			for arr[i] < mid && i < j { // 从左向右找第一个大于等于参照位的数
				i++
			}
			if i <= j {
				// 交换两个数的位置
				swap(arr[i], arr[j])
			}
			// i j 下标的数，已经相对有序了
		}
		// 此时 i j 碰头了
		// 递归操作前一部分
		if start <= i {
			FastSort(arr, start, i-1)
		}
		// 递归操作后一部分
		if end >= j {
			FastSort(arr, j+1, end)
		}
	}
	return arr
}

func FastSort1(nums []int, lo, hi int) {
	// 前序位置
	p := partition(nums, lo, hi)
	FastSort1(nums, lo, p-1)
	FastSort1(nums, p+1, hi)
}

func partition(nums []int, low, high int) int {
	pivot := arr[high] // 选择最后一个元素作为枢纽元素
	left := low
	right := high

	for left != right {
		// 如果当前元素小于或等于枢纽元素，将其交换到左侧
		for left < right && nums[right] < pivot {
			right--
		}
		for left < right && nums[left] > pivot {
			left++
		}
		//找到left比基准大，right比基准小，进行交换
		if left < right {
			swapByIndex(arr, left, right)
		}
	}

	// 第一轮完成，让left和right重合的位置和基准交换，返回基准的位置
	// 重合位置左边的元素，都小于 pivot，右边的元素都大于 pivot
	swapByIndex(arr, low, left)
	// 按 left 和 right 重合的位置做为分界点，对两边进行分治递归
	return left
}

func swapByIndex(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
