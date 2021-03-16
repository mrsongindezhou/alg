package main

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/17 12:28 上午
 * @description：
 * @modified By：
 * @version    ：$
 */


type TreeNode struct {
 Val int
 Left *TreeNode
 Right *TreeNode
}

//type ListNode struct {
//
//}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil && right != nil {
		return right
	} else if left != nil && right == nil {
		return left
	} else {
		return nil
	}
}