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
执行用时分布
2
ms
击败
34.03%
复杂度分析
消耗内存分布
9.45
MB
击败
43.66%

根据回忆写完
*/
func diameterOfBinaryTree(root *TreeNode) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	maxDiameter := 0
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := height(root.Left)
		right := height(root.Right)

		maxDiameter = max(maxDiameter, left+right)
		return 1 + max(left, right)
	}
	height(root)
	return maxDiameter
}
