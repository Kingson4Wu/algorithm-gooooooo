package _02510

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
看答案做的

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
6.11
MB
击败
31.61%
*/
func generateTrees(n int) []*TreeNode {
	var helper func(start, end int) []*TreeNode
	helper = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		var allTrees []*TreeNode
		for i := start; i <= end; i++ {
			leftTrees := helper(start, i-1)
			rightTrees := helper(i+1, end)
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					root := &TreeNode{Val: i, Left: left, Right: right}
					allTrees = append(allTrees, root)
				}
			}
		}
		return allTrees
	}
	return helper(1, n)
}
