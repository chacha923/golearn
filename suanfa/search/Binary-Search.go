package search

//二分查找，返回key值在数组中的下标，否则返回-1
//要求数组有序
//时间复杂度 logn
func binarySearch(array []int, key int) int{
	var left int = 0
	var right int = len(array) - 1

	for left <= right {
		var mid int = (left + right) / 2
		if array[mid] == key {
			return mid
		} else if array[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
