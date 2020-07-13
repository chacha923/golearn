package other

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNumber(nums []int) int {
	ans := nums[0]
	if len(nums) == 1 {
		return ans
	}
	for i := 1; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}
	return ans
}
