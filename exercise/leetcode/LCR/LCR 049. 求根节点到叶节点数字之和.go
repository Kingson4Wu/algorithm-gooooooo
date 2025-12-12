package LCR

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.04
MB
击败
22.86%

还是比较容易，深度优先递归
*/
func sumNumbers(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode, pre int)
	dfs = func(root *TreeNode, pre int) {
		if root == nil {
			return
		}
		pre = pre*10 + root.Val
		if root.Left == nil && root.Right == nil {
			ans += pre
			return
		}
		dfs(root.Left, pre)
		dfs(root.Right, pre)
	}
	dfs(root, 0)
	return ans
}
