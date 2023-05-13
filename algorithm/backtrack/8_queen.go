package backtrack

var solutions [][]string

// 8 皇后，比较有难度 建立回溯思想，一般不会考
// 求 N 个皇后 在 N 维度的棋盘上的所有解
// 皇后的走法：横竖斜都可以走，但是不能走到其他皇后的位置
// 时间复杂度：O(N!)，其中 N 是皇后数量。
// 空间复杂度：O(N)
func solveNQueens(n int) [][]string {
	// 初始化棋盘
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	// 记录每一行皇后的位置，必然的，每行只能有一个皇后
	var queens = make([]int, n)
	// 我们从上到下尝试每一行的皇后放置位置
	// 辅助数组，记录每一列，每一条对角线是否有皇后
	//	方向一的斜线为从左上到右下方向，同一条斜线上的每个位置满足行下标与列下标之差相等，例如 (0,0)(0,0) 和 (3,3)(3,3) 在同一条方向一的斜线上。因此使用行下标与列下标之差即可明确表示每一条方向一的斜线。
	//	方向二的斜线为从右上到左下方向，同一条斜线上的每个位置满足行下标与列下标之和相等，例如 (3,0)(3,0) 和 (1,2)(1,2) 在同一条方向二的斜线上。因此使用行下标与列下标之和即可明确表示每一条方向二的斜线。
	var col, main, sub = make(map[int]bool, n), make(map[int]bool, n), make(map[int]bool, n)
	searchQueen(queens, n, 0, col, main, sub)
	return solutions
}

func searchQueen(queens []int, n int, row int, col, main, sub map[int]bool) {
	// 递归终止条件，访问完所有行了
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	// 遍历每一列
	for i := 0; i < n; i++ {
		// 剪枝
		// 列有皇后
		if col[i] {
			continue
		}
		// 对角线有皇后
		if main[row-i] {
			continue
		}
		if sub[row+i] {
			continue
		}
		// 放皇后
		queens[row] = i
		col[i] = true
		main[row-i] = true
		sub[row+i] = true
		searchQueen(queens, n, row+1, col, main, sub)
		// 回溯
		delete(col, i)
		delete(main, row-i)
		delete(sub, row+i)
	}
}

// 生成棋盘的一个解
func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		var row = make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		// 第 i 行的皇后放在了第 queens[i] 列
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
