package search

// 求算术平方根，实际上是二分查找，递归
// 8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
// x的算术平方根平方后小于等于x
func mySqrt(x int) int {
	var ans = -1
	var l = 0
	var r = x
	for l <= r {
		var mid = l + (r-l)/2 // 防止溢出
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else if mid*mid < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
