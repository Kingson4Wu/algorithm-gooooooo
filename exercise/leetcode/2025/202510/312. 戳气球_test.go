package _02510

import (
	"fmt"
	"testing"
)

func maxCoins(nums []int) int {

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dp := make([][]int, len(nums)+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums)+2)
	}
	val := append([]int{1}, append(nums, 1)...)

	for step := 2; step < len(nums)+2; step++ {
		for i := 0; i+step < len(nums)+2; i++ {
			j := i + step
			//计算时，k处于中间，不能等于i也不能等于j
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], val[i]*val[k]*val[j]+dp[i][k]+dp[k][j])
			}
		}
	}

	return dp[0][len(nums)+1]
}

func TestMaxCoins(t *testing.T) {

	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
	fmt.Println(maxCoins([]int{1, 5}))
}

/**
执行用时分布
36
ms
击败
73.08%
复杂度分析
消耗内存分布
7.38
MB
击败
26.92%
*/
