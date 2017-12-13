package test

//查找匹配的子字符串, 返回子串出现的第一个位置
//优化的方法: kmp

func strStr(source string, target string) int {
	if source == "" && target == "" {
		return -1
	}

	var i, j int
	for i = 0; i < len(source); i++ {
		for j = 0; j < len(target); j++ {
			if source[i+j] != target[j] {
				break
			}
		}
		if j == len(target) {
			return i
		}
	}
	return -1
}
