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
3
ms
击败
93.55%
复杂度分析
消耗内存分布
7.29
MB
击败
90.32%

比较容易做出来
*/
func largestValues(root *TreeNode) []int {
	var ans []int
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(ans) == depth {
			ans = append(ans, root.Val)
		} else if root.Val > ans[depth] {
			ans[depth] = root.Val
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return ans
}
