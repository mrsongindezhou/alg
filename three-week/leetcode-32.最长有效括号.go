package _021_03_28

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/28 11:18 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
func longestValidParentheses(s string) int {
	res := 0
	n := len(s)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' { // ()
				dp[i] = 2
				if i-2 >= 0 {
					dp[i] = dp[i-2] + dp[i]
				}
			} else if dp[i-1] > 0 { // ()((()))
				if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
					dp[i] = dp[i-1] + 2
					if i-2-dp[i-1] >= 0 {
						dp[i] = dp[i] + dp[i-2-dp[i-1]]
					}
				}
			}

		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
