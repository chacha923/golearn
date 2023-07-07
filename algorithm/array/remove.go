package array

// 原地删除目标的元素，双指针法
func removeTarget(nums []int, target int) []int {
	var slow = 0
	var fast = 0

	if fast < len(nums) {
		if nums[fast] != target {
			// 保留，复制到 slow 下标
			nums[slow] = nums[fast]
			slow++
			fast++
		} else if nums[fast] == target {
			// 待删除
			fast++
		}
	}

	return nums[:slow]
}
