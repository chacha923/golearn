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

//堆排序
//s[0]不用，实际元素从角标1开始
//父节点元素大于子节点元素
//左子节点角标为2*k
//右子节点角标为2*k+1
//父节点角标为k/2
func HeapSort2(s []int) {
	n := len(s) - 1
	//构造堆
	//如果给两个已构造好的堆添加一个共同父节点，
	//将新添加的节点作一次下沉将构造一个新堆，
	//由于叶子节点都可看作一个构造好的堆，所以
	//可以从最后一个非叶子节点开始下沉，直至
	//根节点，最后一个非叶子节点是最后一个叶子
	//节点的父节点，角标为N/2
	for k := n / 2; k >= 1; k-- {

	}
}

func sink(s []int, k, n int) {
	for {
		i := 2 * k
	}
}
