package topinterview150

import "strings"

/**
跟这道题完全一模一样的解法
exercise/leetcode/topinterview150/205. 同构字符串_test.go
少了用例限制！

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.77
MB
击败
82.07%

*/
/*
*
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", s = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", s = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false
*/
func wordPattern(pattern string, s string) bool {
	m := make(map[uint8]string)
	rm := make(map[string]struct{})
	arr := strings.Split(s, " ")

	if len(pattern) != len(arr) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if ch, ok := m[pattern[i]]; ok {
			if arr[i] != ch {
				return false
			}
		} else {
			if _, ok := rm[arr[i]]; ok {
				return false
			}
			m[pattern[i]] = arr[i]
			rm[arr[i]] = struct{}{}
		}
	}
	return true
}

/**
解答错误
19 / 44 个通过的测试用例

官方题解
输入
pattern =
"aaa"
s =
"aa aa aa aa"

添加到测试用例
输出
true
预期结果
false
*/
