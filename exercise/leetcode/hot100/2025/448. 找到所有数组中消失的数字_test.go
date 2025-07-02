package _025

/*
*
利用负数
遍历第一遍，将存在的数字对应的下标的数字变成负的
遍历第二遍，统计数字大于0的下标，即不存在的数字
*/
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		index := nums[i]
		if index < 0 {
			index = -index
		}
		if nums[index-1] > 0 {
			nums[index-1] = -nums[index-1]
		}
	}
	var results []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			results = append(results, i+1)
		}
	}
	return results
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
10.30
MB
击败
24.36%
复杂度分析

*/
