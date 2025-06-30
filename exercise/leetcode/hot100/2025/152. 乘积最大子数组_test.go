package _025

import (
	"fmt"
	"testing"
)

/*
*
当前想法：
动态规划，双重循环
d[i][j]表示子数组i..j的乘积
dp[i][j] = dp[i][j-1] * num[j]

此为动态规划暴力解法，另有：
用max_dp[i]表示以第 i 个元素结尾的最大乘积
用min_dp[i]表示以第 i 个元素结尾的最小乘积（因为可能是负的）
max_dp[i] = max( max_dp[i-1] * num[i], min_dp[i] * num[i], num[i])
max_dp[i] = min( max_dp[i-1] * num[i], min_dp[i] * num[i], num[i])
一次遍历接口算出来
只用两个变量即可
*/
// 动态规划暴力法
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
/*func maxProduct(nums []int) int {

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
}*/

func maxProduct(nums []int) int {

	minF, maxF, maxVal := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mi := maxF, minF

		maxF = max(nums[i], max(mx*nums[i], mi*nums[i]))
		minF = min(nums[i], min(mx*nums[i], mi*nums[i]))
		maxVal = max(maxF, maxVal)
	}
	return maxVal
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
5.22
MB
击败
57.26%

*/

/*func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}*/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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
