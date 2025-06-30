package _025

import (
	"fmt"
	"math"
	"testing"
)

/*
*
当前想法：
动态规划，双重循环
d[i][j]表示子数组i..j的乘积
dp[i][j] = dp[i][j-1] * num[j]
*/
func maxProduct(nums []int) int {

	maxVal := math.MinInt32
	dp := make([]int, len(nums))
	for length := 1; length <= len(nums); length++ {
		for i := 0; i < len(nums)-length+1; i++ {
			if length == 1 {
				dp[i] = nums[i]
			} else {
				dp[i] *= nums[i+length-1]
			}
			if dp[i] > maxVal {
				maxVal = dp[i]
			}
		}
	}
	return maxVal
}

func TestMaxProduct(t *testing.T) {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))

	fmt.Println(maxProduct([]int{-2}))
	fmt.Println(maxProduct([]int{-2, -4, 2}))

	fmt.Println(maxProduct([]int{-4, 2}))

	fmt.Println(maxProduct([]int{-4, 2, 1, -1, -1, 2, 4, 1}))
}

/**
6
0
-2
16
2
16
*/

/**
执行用时分布
419
ms
击败
5.07%
复杂度分析
消耗内存分布
5.63
MB
击败
46.59%
*/
