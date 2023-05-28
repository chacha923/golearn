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
	res := [][]int{} // 结果集

	var backtrack func(nums []int, res *[][]int, start int)
	backtrack = func(nums []int, res *[][]int, start int) {
		// 访问到最后一个元素了，把当前数组加入结果集
		if start == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			*res = append(*res, tmp)
			return
		}

		for i := start; i < len(nums); i++ {
			// 做选择，把start元素后面的所有元素依次交换到数组头部
			nums[i], nums[start] = nums[start], nums[i]
			// 继续递归
			backtrack(nums, res, start+1)
			// 撤销选择
			nums[i], nums[start] = nums[start], nums[i]
		}
	}

	backtrack(nums, &res, 0)
	return res
}

// func backtrack(nums []int, used map[int]bool, path []int, res *[][]int) {
// 	if len(nums) == len(path) { // 排列结束, 此处一定要copy slice之后再放入结果集！！！！！！！！
// 		tmp := make([]int, len(nums))
// 		copy(tmp, pathNums)
// 		result = append(result, tmp)
// 		return
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		if !used[i] {
// 			used[i] = true                        // 标记已使用
// 			pathNums = append(pathNums, nums[i])  // 将数字加入路径尾部
// 			backtrack(nums, pathNums, used)       // 继续dfs深搜
// 			pathNums = pathNums[:len(pathNums)-1] // 将尾部数据弹出，回滚操作
// 			used[i] = false                       // 标记未使用
// 		}
// 	}
// }
