package main

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/22 1:34 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

var solutions [][]string
func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	queens := make([]int, n) // 存放n皇后的行坐标
	for i := 0; i<n; i++ {
		queens[i] = -1
	}
	cols, diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}, map[int]bool{}
	backtrack(queens, 0, n, cols, diagonals1, diagonals2)
	return solutions
}

func backtrack(queens []int, row, n int, cols, diagonals1, diagonals2 map[int]bool) {
	if n == row {
		item := generateboard(queens, n)
		solutions = append(solutions, item)
		return
	}

	for i := 0; i<n; i++ {
		if cols[i] {
			continue
		}
		// 左上方斜对角线 因为公式：行 - 列 为常数 左上满足 行坐标与列坐标相加为常数
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}

		// 右上方斜对角线 公式： 行 + 列 为常数
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		queens[row] = i
		cols[i] = true
		diagonals1[diagonal1] = true
		diagonals2[diagonal2] = true
		backtrack(queens, row + 1, n, cols, diagonals1, diagonals2)
		queens[row] = -1
		delete(cols, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateboard(queens []int, n int) []string {
	board := make([]string, n)
	for i := 0; i<n; i++ {
		row := make([]byte, n)
		for j := 0; j<n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board[i] = string(row)
	}
	return board
}