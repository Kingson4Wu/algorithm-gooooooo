package topinterview150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := len(nums) / 2
	root := &TreeNode{
		Val: nums[index],
	}
	root.Left = sortedArrayToBST(nums[0:index])
	root.Right = sortedArrayToBST(nums[index+1:])
	return root
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
5.09
MB
击败
56.53%

*/
