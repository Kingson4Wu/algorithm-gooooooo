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
定义一个辅助递归函数，一定要经过root, 这个函数记得是只计算左边或右边的，不能包含两边！！！！sum := max(root.Val, max(root.Val+left, root.Val+right))
定一个全局变量计算最终要的结果： maxSum = max(sum, max(maxSum, root.Val+left+right))

忘了辅助递归函数只能包含单边的！！！不然就不是一条路径了，会有多个分叉！！！
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
		sum := max(root.Val, max(root.Val+left, root.Val+right))
		maxSum = max(sum, max(maxSum, root.Val+left+right))
		return sum
	}
	maxPathSumWithRoot(root)
	return maxSum
}

/**
执行用时分布
1
ms
击败
18.12%
复杂度分析
消耗内存分布
10.14
MB
击败
85.00%

*/

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
