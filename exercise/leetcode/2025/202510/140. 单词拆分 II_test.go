package _02510

import (
	"fmt"
	"strings"
	"testing"
)

// 组合回溯
/**
在这题中：

你不是在“从一堆元素中选几个组成组合”；

而是要从 字符串的当前位置出发，依次决定从这里能切出哪个合法单词，然后递归到剩余部分。

即：

每一层的“位置”是当前下标 start；

从这个位置开始，你尝试 所有可能的词前缀；

如果某个前缀在 wordDict 中，就递归处理剩余部分。

→ 每层只处理一个位置的切割逻辑，而不是从一个集合中挑选多个元素。

这就是典型的「构造型 DFS」模式。
*/
/*func wordBreak2(s string, wordDict []string) []string {

	m := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		m[word] = true
	}

	var ans []string
	var words []string

	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			sentence := strings.Join(words, " ")
			ans = append(ans, sentence)
			return
		}
		for i := start; i < len(s); i++ {
			step := 1
			for i+step <= len(s) {
				if m[s[i:i+step]] {
					words = append(words, s[i:i+step])
					dfs(i + step)
					words = words[:len(words)-1]
					break
				}
				step++
			}
		}
	}
	dfs(0)
	return ans
}*/

func wordBreak2(s string, wordDict []string) []string {

	m := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		m[word] = true
	}

	var ans []string
	var words []string

	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			sentence := strings.Join(words, " ")
			ans = append(ans, sentence)
			return
		}
		step := 1
		for start+step <= len(s) {
			if m[s[start:start+step]] {
				words = append(words, s[start:start+step])
				dfs(start + step)
				words = words[:len(words)-1]
				//break
			}
			step++
		}
	}
	dfs(0)
	return ans
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.92
MB
击败
53.10%

*/

func TestWordBreak2(t *testing.T) {
	fmt.Println(wordBreak2("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	/*fmt.Println("===")
	fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	fmt.Println("===")
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))*/

}
