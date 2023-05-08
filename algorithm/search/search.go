package search

//二分查找，返回key值在数组中的下标，否则返回-1
//要求数组有序
//时间复杂度 O(log n)
//空间复杂度 O(1)
//凡是能对半分的，空间复杂度都是 logn
func binarySearch(array []int, key int) int {
	var left = 0
	var right = len(array) - 1
	// 滑动，直到 left right 下标碰头
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

// 递归的二分查找
func binarySearch2(array []int, key int) int {
	return binarySearchRecursive(array, 0, len(array)-1, key)
}

// left 左下标，right 右下标，将在 array 的左右下标范围内查找
func binarySearchRecursive(array []int, left, right int, target int) int {
	// 边界条件
	if left > right {
		return -1
	}
	var mid = left + (right-left)/2 // 防止溢出
	if array[mid] == target {
		return mid
	}
	// target 大于中位数，查右边
	if array[mid] < target {
		return binarySearchRecursive(array, mid+1, right, target)
	}
	//  array[mid] > target
	return binarySearchRecursive(array, left, mid-1, target)
}
