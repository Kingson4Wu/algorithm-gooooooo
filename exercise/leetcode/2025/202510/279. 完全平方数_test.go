package _02510

func numSquares(n int) int {

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for k := 1; k <= i; k++ {
			square := k * k
			if square > i {
				break
			}
			if dp[i] == 0 || dp[i-square]+1 < dp[i] {
				dp[i] = dp[i-square] + 1
			}
		}
	}
	return dp[n]
}

/**
执行用时分布
35
ms
击败
47.97%
复杂度分析
消耗内存分布
7.87
MB
击败
32.18%
*/
