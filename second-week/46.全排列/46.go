package main

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/24 7:43 下午
 * @description：
 * @modified By：
 * @version    ：$
 */
func permute(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	if n == 0 {
		return res
	}
	visit := make([]bool, n)
	path := []int{}
	var dfs func(nums, path []int, visit []bool)
	dfs = func(nums, path []int, visit []bool) {
		if len(path) == n {
			temp := make([]int, n)
			copy(temp, path)
			res = append(res, temp)
		}
		for i, _ := range nums {
			if !visit[i] {
				path = append(path, nums[i])
				visit[i] = true
				dfs(nums, path, visit)
				path = path[:len(path)-1]
				visit[i] = false
			}
		}
	}
	dfs(nums, path, visit)
	return res
}
