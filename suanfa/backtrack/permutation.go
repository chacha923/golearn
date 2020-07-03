package backtrack

// for 选择 in 选择列表:
// 做选择
// backtrack(路径, 选择列表)
// 撤销选择

// 全排列
// 所以递推公式 就是数组中每次拿出一个，把它追加到剩余数组的全排列中
// 递归的出口是，数组只有一个数时，把自己返回
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := [][]int{} // 结果集

	for i, num := range nums {
		// 把num从 nums 拿出去 得到tmp
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])
		// sub 是把num 拿出去后，数组中剩余数据的全排列
		sub := permute(tmp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}
	}

}

func backtrack(nums []int, used map[int]bool, path []int, res *[][]int) {
	if len(nums) == len(path) { // 排列结束, 此处一定要copy slice之后再放入结果集！！！！！！！！
		tmp := make([]int, len(nums))
		copy(tmp, pathNums)
		result = append(result, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true                        // 标记已使用
			pathNums = append(pathNums, nums[i])  // 将数字加入路径尾部
			backtrack(nums, pathNums, used)       // 继续dfs深搜
			pathNums = pathNums[:len(pathNums)-1] // 将尾部数据弹出，回滚操作
			used[i] = false                       // 标记未使用
		}
	}
}
