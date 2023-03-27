package hot100

import (
	"fmt"
	"testing"
)

/**
好像不困难？错觉？

*/

/**
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。



示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0


提示：

0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'
*/

/*
*
输入
s =
"()(()"
添加到测试用例
输出
4
预期结果
2
*/
func longestValidParentheses(s string) int {

	leftNum := 0
	currentLength := 0
	longest := 0
	for _, c := range s {
		if c == '(' {
			leftNum++
			continue
		}
		if c == ')' {
			if leftNum > 0 {
				leftNum--
				currentLength++
				if currentLength > longest {
					longest = currentLength
				}
			} else {
				leftNum = 0
				currentLength = 0
			}
		}
	}

	return longest * 2
}

func TestLongestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("()()"))
	fmt.Println(longestValidParentheses(""))
	fmt.Println(longestValidParentheses("((())()"))
	fmt.Println(longestValidParentheses(")()(())())"))
	fmt.Println(longestValidParentheses("()())()((()())))"))
	fmt.Println(longestValidParentheses("()(()"))
}
