package _021_03_28

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/28 10:26 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && '1' <= s[i] && s[i] <= '6') {
			cur += pre
		}
		pre = tmp
	}
	return cur
}
