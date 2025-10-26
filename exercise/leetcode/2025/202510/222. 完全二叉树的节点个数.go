package _02510

import "math"

/*
*

自己回忆写对了， 但感觉实现好像复杂了。。。

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
9.73
MB
击败
19.54%
复杂度分析
*/
func countNodes(root *TreeNode) int {
	leftHeight := func(root *TreeNode) int {
		count := 0
		for root.Left != nil {
			count++
			root = root.Left
		}
		return count
	}
	rightHeight := func(root *TreeNode) int {
		count := 0
		for root.Right != nil {
			count++
			root = root.Right
		}
		return count
	}
	if root == nil {
		return 0
	}
	left := leftHeight(root)
	right := rightHeight(root)
	if left == right {
		return int(math.Pow(2, float64(left+1)) - 1)
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
