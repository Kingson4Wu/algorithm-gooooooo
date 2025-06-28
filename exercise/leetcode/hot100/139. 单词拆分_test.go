package hot100

import (
	"fmt"
	"testing"
)

/**
知道是动态规划，也没想出来

1、构造一个单词的hashmap
2、双重循环查询是否存在
if dp[j] && wordDictSet[s[j:i]]
    dp[i] = true

leetcode/dp/139. 单词拆分.go

时间
0 ms
击败
100%
内存
2 MB
击败
79.43%

*/

/**
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
     注意，你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
*/

func wordBreak(s string, wordDict []string) bool {

	m := make(map[string]bool, len(wordDict))

	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = true
	}

	dp := make([]bool, len(s)+1)
	//空串肯定为true
	dp[0] = true

	//s中以i结尾, 子串是0...j
	for i := 1; i <= len(s); i++ {
		for j := 0; j <= i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
