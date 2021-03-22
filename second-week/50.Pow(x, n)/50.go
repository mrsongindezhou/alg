package _0_Pow_x__n_

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/23 12:20 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return pow(x, n)
	}
	return 1 / pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := pow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
