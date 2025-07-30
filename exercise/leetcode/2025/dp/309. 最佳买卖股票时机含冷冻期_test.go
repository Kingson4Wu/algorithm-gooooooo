package dp

import (
	"fmt"
	"testing"
)

/**
凭回忆写的
*/
/**
dp[i][0] 第i天持有股票的最大收益
dp[i][1] 第i天不持有股票的最大收益
dp[i][2] 第i天不持有股票但处于冷冻期(今天卖出了！！)的最大收益

dp[i][0] = max( dp[i-1][0], dp[i-1][1] - price[i])
dp[i][1] = max( dp[i-1][1], dp[i-1][2])
dp[i][2] = dp[i-1][0] + price[i-1]

*/
/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.32
MB
击败
30.93%
*/
func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][2])
		dp[i][2] = dp[i-1][0] + prices[i]
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

func TestMaxProfit(t *testing.T) {

	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
}

/**
3
0
*/
