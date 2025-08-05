package _025

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSymmetricTree(root *TreeNode) bool {

	if root == nil {
		return true
	}
	var check func(left, right *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil {
			return false
		}
		if right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return check(left.Left, right.Right) && check(left.Right, right.Left)
	}
	return check(root.Left, root.Right)
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.58
MB
击败
84.13%
复
*/
