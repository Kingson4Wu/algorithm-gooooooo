package _02510

import (
	"fmt"
	"testing"
)

/*
*
根据回忆，
二维数组动态规划？
dp[i][j]
答案就是dp[0][n]
从长度1开始遍历

0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.16
MB
击败
5.13%
*/
func wordBreak(s string, wordDict []string) bool {

	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}

	wordMap := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = true
	}
	for length := 1; length <= len(s); length++ {
		for i := 0; i <= len(s)-length; i++ {
			j := i + length - 1
			if wordMap[s[i:j+1]] {
				dp[i][j] = true
				continue
			}
			k := i
			for k < j {
				if dp[i][k] && dp[k+1][j] {
					dp[i][j] = true
					break
				}
				k++
			}
		}
	}

	return dp[0][len(s)-1]
}

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	//fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	//fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
