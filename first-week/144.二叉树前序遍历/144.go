package main

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/13 1:12 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 4)
	return recursion(root, res)
}

func recursion(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = recursion(root.Left, res)
	res = recursion(root.Right, res)
	return res
}