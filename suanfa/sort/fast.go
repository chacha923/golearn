package sort

import "fmt"

var arr = []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}

//
//var arr = []int{0, 0, 0}

func RunFastSort() {
	fmt.Println(arr)
	FastSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

//快速排序 前后指针法
func FastSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	i := start
	j := end
	mid := arr[i] //取第一个元素为mid元素
	for i < j {
		for arr[j] >= mid && i < j { // 从右向左找第一个小于mid的数
			j--
		}
		for arr[i] <= mid && i < j { // 从左向右找第一个大于mid的数
			i++
		}
		if i < j {
			Swap(arr, i, j)
		}
	}

	if i != start {
		Swap(arr, i, start)
	}
	//arr[i] == mid //此时arr[i]有序
	FastSort(arr, start, i-1)
	FastSort(arr, i+1, end)
}

func partition(arr []int, left, right int) {
	key := arr[left]
	begin := left + 1
	end := right
	for begin < end {
		for begin < end && arr[begin] <= key {
			begin++
		}
		for begin < end && arr[right] >= key {
			end--
		}
		if begin < end {
			Swap(arr, begin, end)
		}
	}
}

// 挖坑法
func FastSort2(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			Swap(data, i, tail)
			tail--
		} else {
			Swap(data, i, head)
			head++
			i++
		}
	}
	FastSort2(data[:head])
	FastSort2(data[head+1:])
}
