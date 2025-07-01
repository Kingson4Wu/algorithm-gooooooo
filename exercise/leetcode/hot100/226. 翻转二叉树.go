package hot100

/**
leetcode/tree/226. 翻转二叉树.go
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}
	root.Right, root.Left = root.Left, root.Right
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.08
MB
击败
10.09%
*/
