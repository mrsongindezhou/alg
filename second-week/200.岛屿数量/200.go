package main

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/24 11:30 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
func numIslands(grid [][]byte) int {
	row, col := len(grid), len(grid[0])
	landsNums := 0
	var dfs func(x, y int, grid [][]byte)
	dfs = func(x, y int, grid [][]byte) {
		if x < 0 || y < 0 || x >= row || y >= col || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(x-1, y, grid)
		dfs(x+1, y, grid)
		dfs(x, y-1, grid)
		dfs(x, y+1, grid)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				landsNums++
				dfs(i, j, grid)
			}
		}
	}
	return landsNums
}
