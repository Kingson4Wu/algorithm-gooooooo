package _02511

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
执行用时分布
3
ms
击败
20.82%
复杂度分析
消耗内存分布
6.73
MB
击败
97.11%
复杂度分析

凭记忆完成
*/
func rob(root *TreeNode) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var rob2 func(root *TreeNode) (int, int)
	rob2 = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		leftRob, leftNotRob := rob2(root.Left)
		rightRob, rightNotRob := rob2(root.Right)
		return root.Val + leftNotRob + rightNotRob, max(leftRob, leftNotRob) + max(rightRob, rightNotRob)
	}
	return max(rob2(root))
}
