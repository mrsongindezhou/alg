package _021_03_28

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/28 8:39 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func minPathSum(grid [][]int) int {
	x, y := len(grid), len(grid[0])
	dp := make([][]int, x+1)
	for i := 0; i <= x; i++ {
		dp[i] = make([]int, y+1)
	}

	// res := 0
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if i == 1 {
				dp[i][j] = dp[i][j-1] + grid[i-1][j-1]
			} else if j == 1 {
				dp[i][j] = dp[i-1][j] + grid[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
			}
		}
	}
	return dp[x][y]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
