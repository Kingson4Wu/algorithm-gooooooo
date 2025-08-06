package _025

import (
	"fmt"
	"testing"
)

//看了题解尝试自己写
/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.23
MB
击败
67.26%

*/

func removeInvalidParentheses(s string) []string {

	// 检查是否有效括号
	checkValid := func(s string) bool {
		count := 0
		for _, c := range s {
			if c == '(' {
				count++
			} else if c == ')' {
				if count == 0 {
					return false
				}
				count--
			}
		}
		return count == 0
	}

	//计算最少要删除多少左右括号
	calDeleteCount := func(s string) (int, int) {
		lRemove, rRemove := 0, 0
		for _, c := range s {
			if c == '(' {
				lRemove++
			} else if c == ')' {
				if lRemove > 0 {
					lRemove--
				} else {
					rRemove++
				}
			}
		}
		return lRemove, rRemove
	}

	var ans []string

	var dfs func(s string, start, lRemove, rRemove int)
	dfs = func(s string, start, lRemove, rRemove int) {
		if rRemove == 0 && lRemove == 0 {
			if checkValid(s) {
				ans = append(ans, s)
			}
			return
		}
		if len(s)-start < lRemove+rRemove {
			return
		}
		for i := start; i < len(s); i++ {
			if i > start && s[i] == s[i-1] {
				continue
			}
			if s[i] == '(' && lRemove > 0 {
				dfs(s[:i]+s[i+1:], i, lRemove-1, rRemove)
			}
			if s[i] == ')' && rRemove > 0 {
				dfs(s[:i]+s[i+1:], i, lRemove, rRemove-1)
			}
		}
	}
	lRemove, rRemove := calDeleteCount(s)
	dfs(s, 0, lRemove, rRemove)
	return ans
}

func TestRemoveInvalidParentheses(t *testing.T) {
	fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("(a)())()"))
	fmt.Println(removeInvalidParentheses(")("))
}

/**
[(())() ()()()]
[(a())() (a)()()]
[]
*/
