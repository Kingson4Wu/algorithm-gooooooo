package top100liked

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
自己做的，直接用递归方便

解答错误
158 / 216 个通过的测试用例

官方题解
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Right != nil {
		subSideView := rightSideView(root.Right)
		return append([]int{root.Val}, subSideView...)
	}
	if root.Left != nil {
		subSideView := rightSideView(root.Left)
		return append([]int{root.Val}, subSideView...)
	}
	return []int{root.Val}
}
