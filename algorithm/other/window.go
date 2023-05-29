package other

// 最小覆盖子串
func minWindow(s string, t string) string {
	return ""
}

// 判断 s 中是否存在 t 的排列
func checkInclusion(t string, s string) bool {
	need := make(map[byte]int)   // t 中字符出现的次数
	window := make(map[byte]int) // 窗口中字符出现的次数
	for _, c := range []byte(t) {
		need[c]++
	}

	var left, right, valid int

	for right < len(s) {
		var c = s[right] // 当前新加入窗口的字符
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 循环条件，窗口长度大于等于 t 的长度
		for right-left >= len(t) {
			// 判断左边是否要收缩
			if valid == len(need) {
				// 找到了
				return true
			}
			var d = s[left] // 即将移出窗口的字符
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
