package LCR

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
*
递归，左右都只有0，就删除

感觉这个做法效率很低，结果还是对了。。。

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.03
MB
击败
85.71%
复杂度分析
*/
func pruneTree(root *TreeNode) *TreeNode {

	var isZeroNode func(root *TreeNode) bool
	isZeroNode = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Val != 0 {
			return false
		}
		left := isZeroNode(root.Left)
		right := isZeroNode(root.Right)
		return left && right
	}

	if root == nil {
		return nil
	}

	if root.Val != 0 {
		root.Left = pruneTree(root.Left)
		root.Right = pruneTree(root.Right)
		return root
	}

	left := isZeroNode(root.Left)
	right := isZeroNode(root.Right)
	if left && right {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	return root
}
