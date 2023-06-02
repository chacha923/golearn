package backtrack

// for 选择 in 选择列表:
// 做选择
// backtrack(路径, 选择列表)
// 撤销选择

// 全排列
// 所以递推公式 就是数组中每次拿出一个，把它追加到剩余数组的全排列中
// 递归的出口是，数组只有一个数时，把自己返回
func FullPermute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := [][]int{}                // 结果集
	track := []int{}                // 临时栈，当打满说明得到了一个最终结果
	used := make([]bool, len(nums)) // 记录 nums[i] 元素是否被使用，简化代码

	var backtrack func(nums []int, track []int, used []bool, res *[][]int)
	backtrack = func(nums []int, track []int, used []bool, res *[][]int) {
		// 触发结束条件
		if len(track) == len(nums) {
			// 因为 track 是全局变量，因此需要新建一个数组来存储一份全排列
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			*res = append(*res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// 排除不合法的选择
			if used[i] {
				// 剪枝，避免重复使用同一个数字
				continue
			}
			// 做选择
			track = append(track, nums[i])
			used[i] = true
			// 进入下一层决策树
			backtrack(nums, track, used, res)
			// 取消选择
			track = track[:len(track)-1]
			used[i] = false
		}
	}

	backtrack(nums, track, used, &res)
	return res
}
