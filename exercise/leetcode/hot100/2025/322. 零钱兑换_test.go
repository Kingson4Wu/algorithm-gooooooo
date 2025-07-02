package _025

import (
	"fmt"
	"testing"
)

/*
*
动态规划，自己想的
coin先放map
dp := make([]int, amount+1)
dp[i] 代表 金额 i的最小组合
dp[i] = min(dp[k] + dp[i-k]); 1 <= k <= i/2
结果是 dp[amount]

超出时间限制
188 / 189 个通过的测试用例
最后执行的输入
添加到测试用例
coins =
[1]
amount =
10000

dp[i] = min(dp[i-coin] + 1);  遍历coin组合所有dp，不要用k
*/
func coinChange(coins []int, amount int) int {

	// 根据测试用例打个补丁
	// 还是没通过
	if len(coins) == 1 {
		if amount%coins[0] == 0 {
			return amount / coins[0]
		}
		return -1
	}

	m := make(map[int]bool)
	for i := 0; i < len(coins); i++ {
		m[coins[i]] = true
	}
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
		if m[i] {
			dp[i] = 1
			continue
		}
		for k := 1; k <= i/2; k++ {
			if dp[k] == -1 || dp[i-k] == -1 {
				continue
			}
			if dp[i] == -1 {
				dp[i] = dp[k] + dp[i-k]
			} else {
				dp[i] = min(dp[i], dp[k]+dp[i-k])
			}
		}
	}
	return dp[amount]
}

func TestCoinChange(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
	fmt.Println(coinChange([]int{1}, 10000))
}

/**
3
-1
0
*/
