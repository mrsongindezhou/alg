package main

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/14 2:23 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	stack := make([]*Node, 0, 10)
	res := make([]int, 0, 10)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res
}