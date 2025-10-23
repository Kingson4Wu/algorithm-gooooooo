package _02510

/*
*
根据回忆，
统计count，以及count对应的元素
*/
func majorityElement(nums []int) int {
	pre := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == pre {
			count++
		} else {
			count--
			if count == 0 {
				pre = nums[i]
				count++
			}
		}
	}
	return pre
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
7.96
MB
击败
79.91%

*/
