package dp

// 518. 零钱兑换 II
//
// 给定一个背包容量为 amount，以及一个物品数组 coins（可重复使用），求有多少种不同的方法可以凑出该金额。
// dp[i][w] 容量为w，前i个的不同方法总数
// dp[i][w] = dp[i-1][w] + dp[i-1][w-coin[i]] //选和不选
//修正，可以重复选，应该是dp[i][w] = dp[i-1][w] + dp[i][w-coin[i]] ！！！！！！！！！

/*
*
❌ 你代码的核心错误在这一句：
dp[i][w] = dp[i-1][w] + dp[i-1][w-coins[i-1]]
⛔ 错误原因：
这个递推式是 0-1 背包的做法，表示「第 i 个硬币取 0 次或 1 次」。

但这道题是完全背包问题，每个硬币可以取多次，因此应该是：

dp[i][w] = dp[i-1][w] + dp[i][w-coins[i-1]]
第二项应该是 dp[i][...]，不是 dp[i-1][...]，表示可以继续取第 i 个硬币。
*/
func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i <= len(coins); i++ {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i <= len(coins); i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= len(coins); i++ {
		for w := 1; w <= amount; w++ {
			if w < coins[i-1] {
				dp[i][w] = dp[i-1][w]
			} else {
				//dp[i][w] = dp[i-1][w] + dp[i-1][w-coins[i-1]]
				dp[i][w] = dp[i-1][w] + dp[i][w-coins[i-1]]
			}
		}
	}
	return dp[len(coins)][amount]
}

/**
执行用时分布
13
ms
击败
20.67%
复杂度分析
消耗内存分布
19.25
MB
击败
18.55%
*/
