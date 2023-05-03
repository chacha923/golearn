package sort

import "fmt"

func RunFastSort() {
	fmt.Println(arr)
	fmt.Println(FastSort(arr, 0, len(arr)-1))
}

//快速排序，递归
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
