package _025

import (
	"fmt"
	"testing"
)

/*
*

这次知道要用动态规划，然后自己花了10分钟左右想出来
感觉写得不好，下标控制有点乱
看了答案其实一维dp就行
dp[i]
s中以i结尾, 子串是0...j

---

动态规划
从最小单位 逐步扩展到最终结果
单词列表用map保存，方便后续使用
dp[i][j] 从i到j的子序列是否符合
双重循环
*/
/**
执行用时分布
2
ms
击败
3.45%
复杂度分析
消耗内存分布
4.18
MB
击败
5.72%
复杂度分析

*/
func wordBreak(s string, wordDict []string) bool {

	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	// dp[i][j] 从i到j的子序列是否符合
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		//长度为1的初始化
		dp[i][i] = wordMap[s[i:i+1]]
	}
	for length := 2; length <= len(s); length++ {
		for i := 0; i+length-1 < len(s); i++ {
			j := i + length - 1
			if wordMap[s[i:j+1]] {
				dp[i][j] = true
				continue
			}
			for mid := i; mid < j; mid++ {
				if dp[i][mid] && dp[mid+1][j] {
					dp[i][j] = true
					break
				}
			}
		}
	}
	return dp[0][len(s)-1]
}

/*
*
解答错误
39 / 47 个通过的测试用例

官方题解
输入
s =
"a"
wordDict =
["a"]

添加到测试用例
输出
false
预期结果
true
*/
func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("a", []string{"a"}))
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
