package main

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/13 12:47 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func moveZero(nums []int) {
	for l, r := 0, 0; r < len(nums)-1; r++ {
		if r != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}
