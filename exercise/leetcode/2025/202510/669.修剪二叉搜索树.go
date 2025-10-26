package _02510

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
根据模糊的回忆，稀里糊涂做出来了

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
7.97
MB
击败
26.11%
*/
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	} else if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
