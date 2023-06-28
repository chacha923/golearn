package other

// 一套比较吉利的楼层计数，把包含 '4' 的楼层给跳过了
// 输入一个真实的楼层数，返回吉利的楼层。据说有 O(1) 级别的实现
// e.g. 3 -> 3; 4 -> 5; 13 -> 15
func luckyLevel(realLevel int) int {
	// 最终拿到的楼层号
	var targetLevel int = 1

	for i := 1; i < realLevel; i++ {
		// 每次 +1，碰到需要跳过的楼层，不断往后找
		targetLevel++
		for levelNotLucky(targetLevel) {
			targetLevel++
		}
	}
	return targetLevel
}

// 应当被跳过的楼层
func levelNotLucky(n int) bool {
	for n > 0 {
		if n%10 == 4 {
			return true
		}
	}
	return false
}
