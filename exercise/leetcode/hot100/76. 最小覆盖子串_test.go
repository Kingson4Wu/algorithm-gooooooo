package hot100

import (
	"fmt"
	"testing"
)

/**
看完题解，使用滑动窗口解决
check另外使用一个变量判断，避免每次循环所有字符

时间
24 ms
击败
77.13%
内存
2.9 MB
击败
56.98%


*/

/*
*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
示例 2：

输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
*/
func minWindow(s string, t string) string {

	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	tm := map[byte]int{}
	sm := map[byte]int{}

	var lackC byte
	lack := 0

	check := func() bool {
		if lack != 0 {
			return lack == 2
		}
		for k, tc := range tm {
			if sm[k] < tc {
				return false
			}
		}
		lack = 2
		return true
	}

	for i := 0; i < len(t); i++ {
		tm[t[i]] += 1
	}

	start, end := 0, 0
	minLength := 0
	minStart := 0
	sm[s[0]] = 1
	for start <= end {
		if check() {
			if minLength == 0 || end-start+1 < minLength {
				minLength = end - start + 1
				minStart = start
			}
			sm[s[start]]--

			if tc, ok := tm[s[start]]; ok {
				if tc > sm[s[start]] {
					lackC = s[start]
					lack = 1
				}
			}

			start++
			continue
		}
		end++
		if end >= len(s) {
			break
		}
		sm[s[end]]++

		if s[end] == lackC {
			if tc, ok := tm[s[end]]; ok {
				if tc <= sm[s[end]] {
					lack = 2
				}
			}
		}

	}

	if minLength == 0 {
		return ""
	}
	return s[minStart : minStart+minLength]
}

/**
解答错误
186 / 268 个通过的测试用例

官方题解
输入
s =
"aab"
t =
"aab"

添加到测试用例
输出
"ab"
预期结果
"aab"
*/

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))

	fmt.Println(minWindow("abbaaavvvaaaaavvva", "aa"))
	fmt.Println(minWindow("abbaaavvvaaaaavvva", "a"))
	fmt.Println(minWindow("abbaaavvvaaaaavvva", "ab"))
	fmt.Println(minWindow("abbaaavvvaaaaavvva", "abvaaaaab"))

}
