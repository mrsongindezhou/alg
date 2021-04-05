package fourth_week

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/4/4 1:36 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func climbStairs(n int) int {
	p, q, r := 0, 1, 0
	for i := 1; i <= n; i++ {
		r = p + q
		p = q
		q = r
	}
	return r
}
