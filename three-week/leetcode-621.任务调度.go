package _021_03_28

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/29 1:19 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

// 参考题解：https://leetcode-cn.com/problems/task-scheduler/solution/tian-tong-si-lu-you-tu-kan-wan-jiu-dong-by-mei-jia/
// 参考题解：https://leetcode-cn.com/problems/task-scheduler/solution/ren-wu-diao-du-qi-golangda-bai-100yong-hu-de-jie-f/

func leastInterval(tasks []byte, n int) int {
	chars := [26]int{}
	// res记录返回结果
	res, count := 0, 0
	for i := 0; i < len(tasks); i++ {
		chars[tasks[i]-'A']++
		if res < chars[tasks[i]-'A'] {
			res = chars[tasks[i]-'A']
			count = 1
		} else if chars[tasks[i]-'A'] == res {
			count++
		}
	}
	if n == 0 || ((res-1)*(n+1)+count < len(tasks)) {
		return len(tasks)
	}
	return (res-1)*(n+1) + count
}
