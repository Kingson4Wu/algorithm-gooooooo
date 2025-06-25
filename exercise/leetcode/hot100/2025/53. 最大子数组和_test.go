package _025

import (
	"fmt"
	"testing"
)

func maxSubArray(nums []int) int {

	max := 0
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			dp[i][j] = dp[i][j-1] + nums[j]
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	return max
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}
