package _025

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root2 == nil {
		return root1
	}
	if root1 == nil {
		return root2
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
8.36
MB
击败
98.60%
复杂度分析

*/
