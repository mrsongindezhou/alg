package main

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/14 1:23 下午
 * @description：
 * @modified By：
 * @version    ：$
 */


type Node struct {
    Val int
    Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack := make([]*Node, 0, 10)
	res := make([]int, 0, 10)
	stack = append(stack, root)
	for len(stack) > 0 {
		size := len(stack)
		for size > 0 {
			node := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			res = append(res, node.Val)
			for _, v := range node.Children {
				stack = append(stack, v)
			}
			size--
		}

	}
	reversal(res)
	return res
}

func reversal(nums []int) {
	length := len(nums)
	for i := 0; i < length / 2; i++ {
		nums[i], nums[length-1-i] = nums[length-1-i], nums[i]
	}
}