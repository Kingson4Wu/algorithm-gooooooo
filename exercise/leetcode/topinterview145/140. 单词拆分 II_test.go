package topinterview145

import (
	"fmt"
	"strings"
	"testing"
)

/**
给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。

注意：词典中的同一个单词可能在分段中被重复使用多次。



示例 1：

输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
输出:["cats and dog","cat sand dog"]
示例 2：

输入:s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
解释: 注意你可以重复使用字典中的单词。
示例 3：

输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
输出:[]


提示：

1 <= s.length <= 20
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 10
s 和 wordDict[i] 仅有小写英文字母组成
wordDict 中所有字符串都 不同
*/
/**
看到这个问题，第一反应又是回溯

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.93
MB
击败
71.84%

跟131. 分割回文串_test.go（中等题）一样的套路，为啥是困难题？？？

*/
func wordBreak(s string, wordDict []string) []string {

	words := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		words[word] = true
	}

	var ans []string
	var path []string
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			ans = append(ans, strings.Join(path, " "))
			return
		}
		if start < len(s) {
			step := 0
			for start+step < len(s) {
				end := start + step
				if words[s[start:end+1]] {
					path = append(path, s[start:end+1])
					dfs(end + 1)
					path = path[:len(path)-1]
				}
				step++
			}
		}
	}
	dfs(0)
	return ans
}

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	fmt.Println("===")
	fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	fmt.Println("===")
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))

}
