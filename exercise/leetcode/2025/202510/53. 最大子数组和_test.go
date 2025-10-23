package _02510

import "math"

func maxSubArray(nums []int) int {
	sum := 0
	maxSum := math.MinInt64
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] > sum {
			sum = nums[i]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

/**
执行用时分布
1
ms
击败
33.32%
复杂度分析
消耗内存分布
9.44
MB
击败
75.09%

*/
