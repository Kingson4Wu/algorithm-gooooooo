package hot100

import (
	"fmt"
	"testing"
)

/*
*
leetcode/dp/70. 爬楼梯.go
*/
/*func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}*/
func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	first := 1
	second := 2

	for i := 2; i < n; i++ {
		first, second = second, first+second
	}
	return second
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.84
MB
击败
17.18%
*/

func TestClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))

}
