package array

type NumArray struct {
	// 前缀和数组
	preSum []int
}

/* 输入一个数组，构造前缀和 */
func Constructor(nums []int) NumArray {
	n := len(nums)
	// preSum[0] = 0，便于计算累加和
	preSum := make([]int, n+1)
	// 计算 nums 的累加和
	for i := 1; i < n+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{preSum: preSum}
}

/* 查询闭区间 [left, right] 的累加和 */
func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}
