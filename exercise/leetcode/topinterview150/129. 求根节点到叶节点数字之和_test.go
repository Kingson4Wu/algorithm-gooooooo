package topinterview150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
递归可解，简单

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.00
MB
击败
97.88%
*/
/**
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。
*/
func sumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(int, *TreeNode)
	dfs = func(preSum int, node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			sum += preSum*10 + node.Val
			return
		}
		if node.Left != nil {
			dfs(preSum*10+node.Val, node.Left)
		}
		if node.Right != nil {
			dfs(preSum*10+node.Val, node.Right)
		}
	}
	dfs(0, root)
	return sum
}
