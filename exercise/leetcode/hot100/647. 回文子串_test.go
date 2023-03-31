package hot100

import (
	"fmt"
	"testing"
)

/**
似曾相识， 用动态规划搞定

时间
8 ms
击败
16.69%
内存
4.5 MB
击败
35.24%

*/

/*
*
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

示例 1：

输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：

输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
*/
func countSubstrings(s string) int {

	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	result := 0

	dp := make([][]bool, len(s))
	for i := range s {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		result++
	}

	for length := 2; length <= len(s); length++ {
		for i := 0; i+length-1 < len(s); i++ {
			j := i + length - 1
			if length == 2 {
				if s[i] == s[j] {
					dp[i][j] = true
					result++
				} else {
					dp[i][j] = false
				}
				continue
			}
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
				if dp[i][j] {
					result++
				}
			} else {
				dp[i][j] = false
			}
		}
	}

	return result
}

func TestCountSubstrings(t *testing.T) {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}
