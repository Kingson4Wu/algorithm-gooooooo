package dp

/*
*
题目描述
给定一个容量为 W 的背包，以及 N 个物品，每个物品有：重量 wt[i] 和价值 val[i]
每种物品只能选择一次，求在不超过总容量 W 的前提下，最大可获得的总价值。
*/
// dp[i][w] 前i个物品，容量为w的最大价值
// dp[i][w] = max(dp[i-1][w], dp[i][w-wt[i]]+ val[i])

/*
细节问题或可优化点
1. dp[0][wt[0]] = val[0] 初始化不足
你只在 wt[0] <= W 时赋值 dp[0][wt[0]] = val[0]，但这只是第一个物品只放一件的情况，而完全背包允许多个。
所以建议在初始化第一个物品时也用循环处理：

	for w := 1; w <= W; w++ {
	    if wt[0] <= w {
	        dp[0][w] = dp[0][w-wt[0]] + val[0]
	    }
	}

用 dp := make([][]int, N+1) 会更好。这样就不用管初始化了
dp[0][w] 代表0个物品
*/
func knapsack(W int, wt, val []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	N := len(wt)
	dp := make([][]int, N)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, W+1)
	}
	/*if wt[0] <= W {
		dp[0][wt[0]] = val[0]
	}*/
	for w := 1; w <= W; w++ {
		if wt[0] <= w {
			dp[0][w] = dp[0][w-wt[0]] + val[0]
		}
	}
	for i := 1; i < len(wt); i++ {
		for w := 1; w <= W; w++ {
			if wt[i] > w {
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = max(dp[i-1][w], dp[i][w-wt[i]]+val[i])
			}
		}
	}
	return dp[N-1][W]
}
