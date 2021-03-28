package _021_03_28

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/28 4:09 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

// 参考此题解的视频理解 https://leetcode-cn.com/problems/longest-common-subsequence/solution/1143-zui-chang-gong-gong-zi-xu-lie-by-ming-zhi-sha/

func longestCommonSubsequence(text1 string, text2 string) int {
	txt1Len := len(text1)
	txt2Len := len(text2)
	dp := make([][]int, txt1Len+1)
	for i := 0; i <= txt1Len; i++ {
		dp[i] = make([]int, txt2Len+1)
	}
	for i := 1; i <= txt1Len; i++ {
		for j := 1; j <= txt2Len; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[txt1Len][txt2Len]
}
