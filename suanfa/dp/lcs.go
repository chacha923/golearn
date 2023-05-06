package dp

// // 求最长公共子序列
// // 输入：text1 = "abcde", text2 = "ace"
// // 输出：3
// // 解释：最长公共子序列是 "ace"，它的长度为 3。
// func lcs(x, y string) {
// 	if len(text1) == 0 || len(text2) == 0 {
// 		return 0
// 	}

// 	c := make([][]int, len(text1)+1)
// 	for i := 0; i <= len(text1); i++ {
// 		c[i] = make([]int, len(text2)+1)
// 	}
// 	// 默认0, 不用处理边界
// 	for i := 1; i <= len(text1); i++ {
// 		for j := 1; j <= len(text2); j++ {
// 			if text1[i-1] == text2[j-1] {
// 				c[i][j] = c[i-1][j-1] + 1
// 			} else if text1[i-1] != text2[j-1] {
// 				c[i][j] = lib.Max(c[i-1][j], c[i][j-1])
// 			}
// 		}
// 	}
// 	return c[len(text1)][len(text2)]
// }

// func LCSLength(x []byte, y []byte) int {
// 	m := len(x)
// 	n := len(y)
// 	lookup := make([][]int, m+1)
// 	for i := range lookup {
// 		lookup[i] = make([]int, n+1)
// 	}
// 	var i, j int
// 	for i = 0; i <= m; i++ {
// 		lookup[i][0] = 0
// 	}
// 	for j = 0; j <= n; j++ {
// 		lookup[0][j] = 0
// 	}

// 	for i = 1; i <= m; i++ {
// 		for j = 1; j <= n; j++ {
// 			if x[i-1] == y[j-1] {
// 				lookup[i][j] = lookup[i-1][j-1] + 1
// 			} else {
// 				len1 := lookup[i-1][j]
// 				len2 := lookup[i][j-1]
// 				if len1 > len2 {
// 					lookup[i][j] = len1
// 				} else {
// 					lookup[i][j] = len2
// 				}
// 			}
// 		}
// 	}
// 	return lookup[m][n]
// }
