package other

// 处理数字问题

/*
回文数字
在不使用额外的内存空间的条件下判断一个整数是否是回文。
回文指逆序和正序完全相同。

数据范围：−2^31 ≤ n ≤ 2^31−1
进阶： 空间复杂度 O(1)，时间复杂度 O(len(n))

提示：
负整数可以是回文吗？（比如-1）
如果你在考虑将数字转化为字符串的话，请注意一下不能使用额外空间的限制
你可以将整数翻转。但是，如果你做过题目“反转数字”，你会知道将整数翻转可能会出现溢出的情况，你怎么处理这个问题？
*/
func IsPalindromeInt(x int) bool {
	// 考虑边界条件
	// 1. 负数
	// 2. 可能溢出
	// 3. 如何原地实现？拿到 x 的位数长度，再双指针遍历就行，不需要 tmp 数组

	// 负数不是回文
	if x < 0 {
		return false
	}
	// 个位数是回文
	if x < 10 {
		return true
	}

	// 保存每个位数上的整数
	var tmp []int
	for x > 0 {
		var p = x % 10
		tmp = append(tmp, p)
		x = x / 10
	}

	var start = 0
	var end = len(tmp) - 1
	for start < end {
		if tmp[start] != tmp[end] {
			return false
		}
		start++
		end--
	}
	return true
}
