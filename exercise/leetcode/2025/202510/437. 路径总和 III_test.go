package _02510

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

根据答案做的

执行用时分布
30
ms
击败
5.13%
复杂度分析
消耗内存分布
5.73
MB
击败
100.00%

*/
func pathSum(root *TreeNode, targetSum int) int {

	if root == nil {
		return 0
	}

	var pathSumWithRoot func(root *TreeNode, targetSum int) int
	pathSumWithRoot = func(root *TreeNode, targetSum int) int {
		if root == nil {
			return 0
		}
		count := 0
		if root.Val == targetSum {
			count = 1
		}
		return count + pathSumWithRoot(root.Left, targetSum-root.Val) + pathSumWithRoot(root.Right, targetSum-root.Val)
	}
	return pathSumWithRoot(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func TestPathSum(t *testing.T) {
	//[10,5,-3,3,2,null,11,3,-2,null,1]
	//fmt.Println(BuildTree([]int{10,5,-3,3,2,nil,11,3,-2,nil,1}))

	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(pathSum(root2, 2))
}
