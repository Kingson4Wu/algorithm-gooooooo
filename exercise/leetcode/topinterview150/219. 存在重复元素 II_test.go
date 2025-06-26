package topinterview150

/**
执行用时分布
27
ms
击败
70.92%
复杂度分析
消耗内存分布
13.18
MB
击败
55.90%
复杂度分析

*/
/*
*
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。

示例 1：

输入：nums = [1,2,3,1], k = 3
输出：true
示例 2：

输入：nums = [1,0,1,1], k = 1
输出：true
示例 3：

输入：nums = [1,2,3,1,2,3], k = 2
输出：false
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if index, ok := m[num]; ok {
			if i-index <= k {
				return true
			}
		}
		m[num] = i
	}
	return false
}
