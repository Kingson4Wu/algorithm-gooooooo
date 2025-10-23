package _02510

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var check func(left, right *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left == nil && right != nil {
			return false
		} else if left != nil && right == nil {
			return false
		} else if left.Val != right.Val {
			return false
		}
		return check(left.Left, right.Right) && check(left.Right, right.Left)
	}
	return check(root.Left, root.Right)
}

/**
看了答案做的！！
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.68
MB
击败
19.11%
*/
