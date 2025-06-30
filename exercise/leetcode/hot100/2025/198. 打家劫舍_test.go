package _025

import (
	"fmt"
	"testing"
)

/*
*
自己想的
dp[i] 是 i房屋一定要偷的最大值
dp[i] = max(nums[i]+ dp[i-2], nums[i]+ dp[i-3])
*/
func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	if len(nums) == 3 {
		return max(nums[0]+nums[2], nums[1])
	}
	dp0, dp1, dp2 := nums[0], nums[1], nums[0]+nums[2]

	maxVal := max(dp0, max(dp1, dp2))
	for i := 3; i < len(nums); i++ {
		dp := max(nums[i]+dp1, nums[i]+dp0)
		maxVal = max(dp, maxVal)
		dp2, dp1, dp0 = dp, dp2, dp1
	}
	return maxVal
}

/**
写得有点丑，虽然这次自己想的
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.88
MB
击败
32.33%
*/

func TestRob(t *testing.T) {
	fmt.Println(rob([]int{2, 1}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

/**
2
4
12
*/
