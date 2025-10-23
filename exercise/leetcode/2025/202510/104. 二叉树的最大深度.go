package _02510

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	max := 0
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if depth > max {
			max = depth
		}
		if root != nil {
			dfs(root.Left, depth+1)
			dfs(root.Right, depth+1)
		}
	}
	dfs(root, 0)
	return max
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
32.44%
*/
