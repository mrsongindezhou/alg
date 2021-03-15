package 剑指Offer_49_丑数

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/16 12:25 上午
 * @description：
 * @modified By：
 * @version    ：$
 */
func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(min(n2, n3), n5)
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
