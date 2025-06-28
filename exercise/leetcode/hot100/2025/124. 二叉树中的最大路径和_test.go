package _025

import "math"

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
看了题解写得
递归只选左或右的函数，设置一个全局变量，在递归函数内部计算root加左右的最大值

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
10.16
MB
击败
60.45%
*/
func maxPathSum(root *TreeNode) int {

	maxSum := math.MinInt32

	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := maxGain(node.Left)
		right := maxGain(node.Right)

		sum := max(node.Val, max(node.Val+left+right, max(node.Val+left, node.Val+right)))
		maxSum = max(sum, maxSum)

		//只选左或右
		return max(node.Val, max(node.Val+left, node.Val+right))
	}
	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
