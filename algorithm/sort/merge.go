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
func mergeSort(arr, tmp []int, left, right int) {
	if left < right {
		// mid := (left + right) / 2
		mid := left + (right-left)/2
		mergeSort(arr, tmp, left, mid)    //左边归并排序，使得左子序列有序
		mergeSort(arr, tmp, mid+1, right) //右边归并排序，使得右子序列有序
		merge(arr, tmp, left, mid, right) //将两个有序子数组合并操作
	}
}

// 将有二个有序数列a[first...mid]和a[mid...last]合并
func merge(arr, tmp []int, left, mid, right int) {
	i := left    //左序列指针
	j := mid + 1 //右序列指针
	t := 0       //临时数组指针
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			tmp[t] = arr[i]
			t++
			i++
		} else {
			tmp[t] = arr[j]
			t++
			j++
		}
	}
	for i <= mid { //将左边剩余元素填充进temp中
		tmp[t] = arr[i]
		t++
		i++
	}
	for j <= right { //将右序列剩余元素填充进temp中
		tmp[t] = arr[j]
		t++
		j++
	}

	t = 0
	//将temp中的元素全部拷贝到原数组中
	for left <= right {
		arr[left] = tmp[t]
		left++
		t++
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
