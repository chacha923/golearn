package search

//二分查找，返回key值在数组中的下标，否则返回-1
//要求数组有序
//时间复杂度 logn
func binarySearch(array []int, key int) int {
	var left = 0
	var right = len(array) - 1

	for left <= right {
		// Prevent (left + right) overflow
		// var mid = (left + right) / 2
		var mid = left + (right-left)/2 // 防止溢出
		if array[mid] == key {
			return mid
		} else if array[mid] < key {
			left = mid + 1
		} else if array[mid] > key {
			right = mid - 1
		}
	}
	// End Condition: left > right
	return -1
}

func binarySearch2(array []int, left, right int, target int) int {
	mid := (left + right) / 2
	if left <= right {
		if array[mid] == target {
			return mid
		} else if target < array[mid] {
			return binarySearch2(array, left, mid-1, target)
		} else {
			return binarySearch2(array, mid+1, right, target)
		}
	}
	return -1
}
