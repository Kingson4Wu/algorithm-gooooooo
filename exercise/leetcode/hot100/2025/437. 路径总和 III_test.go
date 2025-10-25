package _025

import (
	"fmt"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**

时隔两年，没做出来！！！。。

以下写法是错的
*/
func pathSum(root *TreeNode, targetSum int) int {
	count := 0
	var dfs func(root *TreeNode, targetSum int)
	dfs = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}
		if root.Val == targetSum {
			count++
		}
		dfs(root.Left, targetSum-root.Val)
		dfs(root.Left, targetSum)
		dfs(root.Right, targetSum-root.Val)
		dfs(root.Right, targetSum)
	}
	dfs(root, targetSum)
	return count
}

func TestPathSum(t *testing.T) {
	//[10,5,-3,3,2,null,11,3,-2,null,1]
	//fmt.Println(BuildTree([]int{10,5,-3,3,2,nil,11,3,-2,nil,1}))

	/*root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(pathSum(root2, 2))*/
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
