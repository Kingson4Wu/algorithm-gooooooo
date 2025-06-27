package topinterview150

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
不想用栈写，用递归
超出内存限制

改用队列写
*/

/*
*
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
*/
/*func zigzagLevelOrder(root *TreeNode) [][]int {
	var results [][]int
	var dfs func(int, *TreeNode)
	dfs = func(i int, node *TreeNode) {
		if node == nil {
			return
		}
		if len(results) == i {
			results = append(results, []int{node.Val})
		} else {
			if i%2 == 0 {
				results[i] = append(results[i], node.Val)
			} else {
				results[i] = append(append([]int{}, node.Val), results[i]...)
			}
		}
		if i%2 == 0 {
			dfs(i+1, root.Left)
			dfs(i+1, root.Right)
		} else {
			dfs(i+1, root.Right)
			dfs(i+1, root.Left)
		}
	}
	dfs(0, root)
	return results
}*/
/**
超出内存限制
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var results [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	level := 0

	for len(queue) > 0 {
		var result []int
		var newQueue []*TreeNode
		if level%2 == 0 {
			for i := 0; i < len(queue); i++ {
				result = append(result, queue[i].Val)
				if queue[i].Left != nil {
					newQueue = append(newQueue, queue[i].Left)
				}
				if queue[i].Right != nil {
					newQueue = append(newQueue, queue[i].Right)
				}
			}
		} else {
			for i := len(queue) - 1; i >= 0; i-- {
				result = append(result, queue[i].Val)
				if queue[i].Right != nil {
					newQueue = append(append([]*TreeNode{}, queue[i].Right), newQueue...)
				}
				if queue[i].Left != nil {
					newQueue = append(append([]*TreeNode{}, queue[i].Left), newQueue...)
				}
			}
		}

		results = append(results, result)
		queue = newQueue
		level++
	}

	return results
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.70
MB
击败
5.69%
复杂度分析

*/
