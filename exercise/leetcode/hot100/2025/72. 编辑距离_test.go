package _025

import (
	"fmt"
	"testing"
)

/*
*
凭回忆写的
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				//插入、删除、替换
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}

func TestMinDistance(t *testing.T) {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}

/**
3
5
*/

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
7.22
MB
击败
45.54%
复杂度分析

*/
