package _02510

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
根据回忆写
定义一个辅助递归函数，一定要经过root
定一个全局变量计算最终要的结果
*/
func maxPathSum(root *TreeNode) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	maxSum := math.MinInt32
	var maxPathSumWithRoot func(root *TreeNode) int
	maxPathSumWithRoot = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := maxPathSumWithRoot(root.Left)
		right := maxPathSumWithRoot(root.Right)
		sum := max(root.Val, max(root.Val+left, max(root.Val+right, root.Val+left+right)))
		maxSum = max(sum, maxSum)
		return sum
	}
	maxPathSumWithRoot(root)
	return maxSum
}

/**

输入
root =
[5,4,8,11,null,13,4,7,2,null,null,null,1]

添加到测试用例
输出
55
预期结果
48
*/
