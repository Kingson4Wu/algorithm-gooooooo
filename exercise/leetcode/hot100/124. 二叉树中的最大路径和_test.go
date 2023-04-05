package hot100

import (
	"fmt"
	"testing"
)

/**
递归写完了。。。
没考虑根结点是负数的情况。。
*/

/**
输入
root =
[-3]
添加到测试用例
输出
0
预期结果
-3

输入
root =
[1,-2,-3,1,3,-2,null,-1]
添加到测试用例
输出
4
预期结果
3
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/** 这个写法是错的，没考虑要相连 */
func maxPathSum(root *TreeNode) int {

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	if root.Left == nil {
		right := maxPathSum(root.Right)

		maxVal := root.Val

		if right > maxVal {
			maxVal = right
		}

		if right+root.Val > maxVal {
			maxVal = right + root.Val
		}
		return maxVal
	}
	if root.Right == nil {
		left := maxPathSum(root.Left)

		maxVal := root.Val

		if left > maxVal {
			maxVal = left
		}

		if left+root.Val > maxVal {
			maxVal = left + root.Val
		}
		return maxVal
	}

	left := maxPathSum(root.Left)
	right := maxPathSum(root.Right)

	maxVal := root.Val

	if left > maxVal {
		maxVal = left
	}
	if right > maxVal {
		maxVal = right
	}

	if left+root.Val > maxVal {
		maxVal = left + root.Val
	}
	if right+root.Val > maxVal {
		maxVal = right + root.Val
	}
	if left+right+root.Val > maxVal {
		maxVal = left + right + root.Val
	}

	return maxVal
}

func TestMaxPathSum(t *testing.T) {
	//[1,-2,-3,1,3,-2,null,-1]
	fmt.Println(maxPathSum(&TreeNode{Val: 1, Left: &TreeNode{Val: -2, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: -1}}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: -3, Left: &TreeNode{Val: -2}}}))

}
