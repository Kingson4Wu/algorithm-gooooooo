package hot100

import (
	"fmt"
	"testing"
)

/**
leetcode/tree/437. 路径总和 III.go
*/

/**
自己用递归，一直没对
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

	return pathSum2(root, targetSum, targetSum)
}

func pathSum2(root *TreeNode, totalSum, targetSum int) int {

	if root == nil {
		return 0
	}

	count := 0
	if root.Val == targetSum {
		count = 1
	}
	return count + pathSum2(root.Left, totalSum, totalSum) + pathSum2(root.Right, totalSum, totalSum) + pathSum2(root.Left, totalSum, targetSum-root.Val) + pathSum2(root.Right, totalSum, targetSum-root.Val)
}

func TestPathSum(t *testing.T) {
	//[10,5,-3,3,2,null,11,3,-2,null,1]
	//fmt.Println(BuildTree([]int{10,5,-3,3,2,nil,11,3,-2,nil,1}))

	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(pathSum(root2, 2))
	//1

	root := &TreeNode{Val: 10, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}}}

	fmt.Println(pathSum(root, 8))
	//3

	/**
	解答错误

	100 / 128 个通过的测试用例
	输入
	root =
	[1,null,2,null,3,null,4,null,5]
	targetSum =
	3
	添加到测试用例
	输出
	3
	预期结果
	2
	*/
}
