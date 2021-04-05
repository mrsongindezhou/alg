package fourth_week

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/4 1:37 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func numIslands(grid [][]byte) int {
	x, y := len(grid), len(grid[0])
	dp := make([][]byte, x)
	landsNums := 0
	for i := 0; i < x; i++ {
		dp[i] = make([]byte, y)
	}
	var dfs func(i, j int, grid [][]byte)
	dfs = func(i, j int, grid [][]byte) {
		if i < 0 || j < 0 || i >= x || j >= y || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i-1, j, grid)
		dfs(i+1, j, grid)
		dfs(i, j-1, grid)
		dfs(i, j+1, grid)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if grid[i][j] == '1' {
				landsNums++
				dfs(i, j, grid)
			}
		}
	}
	return landsNums
}
