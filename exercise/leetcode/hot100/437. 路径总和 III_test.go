package hot100

import (
	"fmt"
	"testing"
)

/**
leetcode/tree/437. 路径总和 III.go
*/

/**
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。


*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {

	if root == nil {
		return 0
	}

	count := 0
	if root.Val == targetSum {
		count = 1
	}
	return count + pathSum(root.Left, 8) + pathSum(root.Right, 8) + pathSum(root.Left, targetSum-root.Val) + pathSum(root.Right, targetSum-root.Val)
}

func TestPathSum(t *testing.T) {
	//[10,5,-3,3,2,null,11,3,-2,null,1]
	//fmt.Println(BuildTree([]int{10,5,-3,3,2,nil,11,3,-2,nil,1}))

	root := &TreeNode{Val: 10, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}}}

	fmt.Println(pathSum(root, 8))
}
