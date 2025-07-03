package _025

/*
*
动态规划
if s[i] == s[j]
dp[i][j] = dp[i+1][j-1]
*/
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	count := n
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				if length == 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] {
					count++
				}
			}
		}
	}
	return count
}

/**
执行用时分布
8
ms
击败
14.73%
复杂度分析
消耗内存分布
7.92
MB
击败
16.18%
*/
