package sort

import "fmt"

var arr = []int{100, 16, 4, 8, 70, 2, 37, 23, 5, 12}

func RunFastSort() {
	fmt.Println(FastSort(arr, 0, len(arr)-1))
}

//快速排序
func FastSort(arr []int, start int, end int) []int {
	if start < end {
		i := start
		j := end
		mid := arr[i]
		for i < j {
			for arr[j] > mid && i < j {	// 从右向左找第一个小于x的数
				j--
			}
			for arr[i] < mid && i < j {	// 从左向右找第一个大于等于x的数
				i++
			}
			if i <= j {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}

		}
		arr[i] = mid
		if start <= i {
			FastSort(arr, start, i-1)
		}
		if end >= j {
			FastSort(arr, j+1, end)
		}
	}
	return arr
}
