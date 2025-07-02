package hot100

/**
leetcode/tree/538. 把二叉搜索树转换为累加树.go
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
自己想的
用中序的倒序来递归遍历累加

执行用时分布
4
ms
击败
3.05%
复杂度分析
消耗内存分布
8.50
MB
击败
54.57%

*/
func convertBST(root *TreeNode) *TreeNode {
	pre := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Right != nil {
			dfs(root.Right)
		}
		root.Val += pre
		pre = root.Val
		if root.Left != nil {
			dfs(root.Left)
		}
	}
	dfs(root)
	return root
}
