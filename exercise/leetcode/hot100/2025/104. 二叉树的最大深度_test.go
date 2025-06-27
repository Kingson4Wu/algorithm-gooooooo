package _025

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	var dfs func(depth int, root *TreeNode) int
	dfs = func(depth int, root *TreeNode) int {
		if root.Left == nil && root.Right == nil {
			return depth + 1
		} else if root.Left != nil && root.Right != nil {
			left := dfs(depth+1, root.Left)
			right := dfs(depth+1, root.Right)
			if left > right {
				return left
			}
			return right
		} else if root.Left != nil {
			return dfs(depth+1, root.Left)
		}
		return dfs(depth+1, root.Right)
	}

	return dfs(0, root)
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
6.18
MB
击败
32.41%
*/
