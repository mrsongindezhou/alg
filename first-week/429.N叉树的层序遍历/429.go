package main

/**
 * @author     ：songdehua
 * @emall      ：200637086@qq.com
 * @date       ：Created in 2021/3/14 12:14 下午
 * @description：
 * @modified By：
 * @version    ：$
 */

type Node struct {
	Val int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 声明一个队列
	queue := make([]*Node, 0, 10)
	// 声明结果集，用于保存结果
	res := make([][]int, 0, 10)
	// 将root节点放进队列
	queue = append(queue, root)
	// BFS搜索
	for len(queue) > 0 {
		// 记录当前队列状态的大小
		size := len(queue)
		items := make([]int, size)
		for i := 0; i < size; i++ {
			// 取出队列头元素
			node := queue[0]
			queue = queue[1:]
			// 放到结果表
			items[i] = node.Val
			for _, v := range node.Children {
				queue = append(queue, v)
			}
		}
		res = append(res, items)
	}
	return res
}