package _02510

import (
	"fmt"
	"strings"
	"testing"
)

// 组合回溯
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
}

func TestWordBreak2(t *testing.T) {
	fmt.Println(wordBreak2("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	/*fmt.Println("===")
	fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	fmt.Println("===")
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))*/

}
