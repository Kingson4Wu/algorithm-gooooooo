package topinterview150

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
用栈可以实现，但我不想用栈
用递归方便

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
8.57
MB
击败
5.07%
*/
/**
给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。
*/
func averageOfLevels(root *TreeNode) []float64 {

	nums := make(map[int]int)
	sums := make(map[int]int)

	var dfs func(int, *TreeNode)
	dfs = func(i int, root *TreeNode) {
		nums[i]++
		sums[i] += root.Val
		if root.Left != nil {
			dfs(i+1, root.Left)
		}
		if root.Right != nil {
			dfs(i+1, root.Right)
		}
	}
	dfs(0, root)
	result := make([]float64, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = float64(sums[i]) / float64(nums[i])
	}

	return result
}
