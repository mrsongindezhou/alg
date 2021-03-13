package _2_接雨水

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/13 3:55 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func trap(height []int) int {
	// 数组长度
	lenght := len(height)
	if lenght == 0 {
		return 0
	}
	// 声明变量记录答案
	var ans int
	maxLeft := make([]int, lenght, lenght)
	maxRight := make([]int, lenght, lenght)

	// 左边最大值
	maxLeft[0] = height[0]
	for i := 1; i<lenght; i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
	}

	// 右边最大值
	maxRight[lenght-1] = height[lenght-1]
	for i := lenght - 2; i >= 0; i-- { // i要大于0
		maxRight[i] = max(height[i], maxRight[i+1])
	}

	for i := 1; i<lenght-1; i++ {
		ans += min(maxLeft[i], maxRight[i]) - height[i]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}