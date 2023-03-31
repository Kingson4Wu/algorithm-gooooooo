package hot100

import (
	"fmt"
	"testing"
)

/**
dp[i][j] = d[i][j-1] + d[j][j]
*/
/*
*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。

示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2
*/
func subarraySum(nums []int, k int) int {

	if len(nums) == 1 {
		if nums[0] == k {
			return 1
		}
		return 0
	}

	result := 0

	dp := make([][]int, len(nums))
	for i := range nums {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
		if nums[i] == k {
			result++
		}
	}

	for length := 2; length <= len(nums); length++ {
		for i := 0; i+length-1 < len(nums); i++ {
			j := i + length - 1
			dp[i][j] = dp[i][j-1] + dp[j][j]
			if dp[i][j] == k {
				result++
			}
		}
	}

	return result
}

func TestSubarraySum(t *testing.T) {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
