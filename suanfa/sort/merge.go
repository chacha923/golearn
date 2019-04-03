package sort

import "fmt"

func RunMergeSort() {
	tmp := make([]int, len(arr))
	mergeSort(arr, tmp, 0, len(arr)-1)
	fmt.Println(arr)
}

func mergeSort(arr, tmp []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(arr, tmp, left, mid)
		mergeSort(arr, tmp, mid+1, right)
		merge(arr, tmp, left, mid, right)
	}
}

//将有二个有序数列a[first...mid]和a[mid...last]合并
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
