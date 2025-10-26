package _02510

import "sort"

/*
*
根据回忆，按最右排
想象往最右射一箭

执行用时分布
55
ms
击败
75.11%
复杂度分析
消耗内存分布
18.71
MB
击败
31.79%
复杂度分析
*/
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1
	maxRight := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > maxRight {
			maxRight = points[i][1]
			count++
		}
	}
	return count
}
