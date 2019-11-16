package sort

import "fmt"

var arr = []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}

func RunFastSort() {
	fmt.Println(arr)
	fmt.Println(FastSort(arr, 0, len(arr)-1))
}

//快速排序
func FastSort(arr []int, start int, end int) []int {
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
			//fmt.Println(arr)
		}
		//arr[i] = mid	//此时arr[i]有序

		if start <= i {
			FastSort(arr, start, i-1)
		}
		if end >= j {
			FastSort(arr, j+1, end)
		}
	}
	return arr
}
