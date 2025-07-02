package _025

import "testing"

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
递归，传root，返回偷和不偷当前结点的最大值
维护一个全局变量计算最大值
*/

func rob3(root *TreeNode) int {
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		robL, notRobL := dfs(root.Left)
		robR, notRobR := dfs(root.Right)
		return root.Val + notRobL + notRobR, max(robL, notRobL) + max(robR, notRobR)
	}
	return max(dfs(root))
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
6.79
MB
击败
89.67%

*/

func TestRob3(t *testing.T) {

}
