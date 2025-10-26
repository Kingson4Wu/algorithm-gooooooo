package _02510

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
根据回忆完成

执行用时分布
6
ms
击败
45.45%
复杂度分析
消耗内存分布
6.93
MB
击败
84.85%

*/
func findBottomLeftValue(root *TreeNode) int {
	v := root.Val
	height := 0
	var dfs func(root *TreeNode, h int)
	dfs = func(root *TreeNode, h int) {
		if root == nil {
			return
		}
		if h+1 > height {
			height = h + 1
			v = root.Val
		}
		dfs(root.Left, h+1)
		dfs(root.Right, h+1)
	}
	dfs(root, 0)
	return v
}
