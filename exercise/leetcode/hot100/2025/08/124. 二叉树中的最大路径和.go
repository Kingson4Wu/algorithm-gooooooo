package _8

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
回忆中写对

执行用时分布
4
ms
击败
100.00%
复杂度分析
消耗内存分布
10.10
MB
击败
58.82%
复杂度分析
*/
func maxPathSum(root *TreeNode) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	maxSum := math.MinInt32
	var maxSinglePathSumWithRoot func(root *TreeNode) int
	maxSinglePathSumWithRoot = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := maxSinglePathSumWithRoot(root.Left)
		right := maxSinglePathSumWithRoot(root.Right)
		maxSingleSum := max(root.Val, max(left+root.Val, right+root.Val))
		maxSum = max(maxSum, max(maxSingleSum, left+right+root.Val))
		return maxSingleSum
	}
	maxSinglePathSumWithRoot(root)
	return maxSum
}
