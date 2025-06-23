package hot100

import (
	"fmt"
	"testing"
)

/**
给你一个字符串 s，找到 s 中最长的 回文 子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
*/
/**
看题解之后做的
*/
/*
*
用 P(i,j) 表示字符串 s 的第 i 到 j 个字母组成的串（下文表示成 s[i:j]）是否为回文串
P(i, j) = P(i+1, j-1)  (S_i == S_j)
也就是说，只有 s[i+1:j-1] 是回文串，并且 s 的第 i 和 j 个字母相同时，s[i:j] 才会是回文串。
*/
/**
1. 初始化长度为1的为true
2. 双重循环，第一个循环是子串的长度，从2开始； 第二个循环是子串开始下标
*/
/**

时间
68 ms
击败
28.71%
内存
7.2 MB
击败
8.53%
*/
func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}

	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	for i := 0; i < length; i++ {
		dp[i][i] = true
	}
	maxStart := 0
	maxLen := 1
	for subLen := 2; subLen <= length; subLen++ {
		for start := 0; start < length; start++ {
			end := start + subLen - 1
			if end >= length {
				continue
			}
			if s[start] != s[end] {
				dp[start][end] = false
				continue
			}
			if subLen == 2 {
				dp[start][end] = true
			} else {
				dp[start][end] = dp[start+1][end-1]
			}

			if dp[start][end] {
				if end-start+1 > maxLen {
					maxLen = end - start + 1
					maxStart = start
				}
			}
		}
	}
	return s[maxStart : maxStart+maxLen]
}

func TestLongestPalindrome(t *testing.T) {

	fmt.Println(longestPalindrome("abcabcbb"))
	fmt.Println(longestPalindrome("bbbbb"))
	fmt.Println(longestPalindrome("pwwkew"))
	fmt.Println(longestPalindrome(" "))

}
