package _02510

import "math"

func maxSubarraySumCircular(nums []int) int {

	maxSubArray := func(nums []int) int {
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
	minSubArray := func(nums []int) int {
		sum := 0
		minSum := math.MaxInt64
		for i := 0; i < len(nums); i++ {
			sum += nums[i]
			if nums[i] < sum {
				sum = nums[i]
			}
			if sum < minSum {
				minSum = sum
			}
		}
		return minSum
	}
	maxSum := maxSubArray(nums)
	if maxSum < 0 {
		return maxSum
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	minSum := minSubArray(nums)
	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}
	return max(maxSum, totalSum-minSum)
}

/**
看完答案自己写
执行用时分布
2
ms
击败
30.18%
复杂度分析
消耗内存分布
8.64
MB
击败
57.48%

*/
