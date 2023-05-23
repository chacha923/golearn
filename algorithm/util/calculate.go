package util

func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func Less(x, y int) bool {
	if x <= y {
		return true
	}
	return false
}

// 求中间下标，防止溢出
func MidIndex(left, right int) int {
	// return (left + right) / 2  => left + right/2 - left/2
	return left + (right-left)/2
}

func Swap(x, y *int) {
	*x, *y = *y, *x
}