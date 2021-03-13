package main

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/13 9:13 上午
 * @description：
 * @modified By：
 * @version    ：$
 */

// desc: 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数

// 第一种方法 旋转后的位置 = (k(移动偏移量) + i(当前位置)) % length(数组长度)
//func rotate(nums []int, k int)  {
//	lenght := len(nums)
//	arr := make([]int, lenght)
//	for i, v := range nums {
//		arr[(i + k) % lenght] = v
//	}
//	copy(nums, arr)
//}

// 方法2 递归实现反转数组
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k]) // 此处是k，因为左闭右开，正好是k个数
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
