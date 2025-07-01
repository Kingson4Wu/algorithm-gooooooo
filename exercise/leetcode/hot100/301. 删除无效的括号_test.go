package hot100

import (
	"fmt"
	"testing"
)

/*
*
给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。

返回所有可能的结果。答案可以按 任意顺序 返回。

示例 1：

输入：s = "()())()"
输出：["(())()","()()()"]
示例 2：

输入：s = "(a)())()"
输出：["(a())()","(a)()()"]
示例 3：

输入：s = ")("
输出：[""]

提示：

1 <= s.length <= 25
s 由小写英文字母以及括号 '(' 和 ')' 组成
s 中至多含 20 个括号
*/
/**
看了题解之后做

方法一：回溯 + 剪枝
1、定义一个函数判断最终的字符串是否是有效括号的字符串
2、先计算总共最少需要删除的左括号和右括号分别的个数（题目就是要求删除最少）
3、深度遍历加回溯，删除左右括号，然后判断是否是有效括号的字符串
4、相同的左括号或右括号（连续），注意剪枝

*/
func removeInvalidParentheses(s string) []string {

	// 判断是不是有效的括号对称字符串
	invalid := func(s string) bool {
		count := 0
		for _, ch := range s {
			if ch == '(' {
				count++
			} else if ch == ')' {
				count--
				if count < 0 {
					//右括号先出现，说明一定不是有效的
					return false
				}
			}
		}
		return count == 0
	}

	var ans []string

	//回溯遍历删除括号（删除最小数量的无效括号），然后判断是否有效的括号对称字符串
	//每次dfs只删除一个字符，只是遍历多个位置删
	var dfs func(s string, start, lRemove, rRemove int)
	dfs = func(s string, start, lRemove, rRemove int) {

		//删够字符了，且是合法字符串
		if lRemove == 0 && rRemove == 0 && invalid(s) {
			ans = append(ans, s)
			return
		}

		// 当前遍历只删除一个字符
		for i := start; i < len(s); i++ {
			//剪枝，发现连续的字符，跳过；不只是开始的连续重复，中间的连续重复字符也能剪枝
			if i != start && s[i] == s[i-1] {
				continue
			}
			// 如果剩余的字符无法满足去掉的数量要求，直接返回
			// 遍历到后面的字串就没必要再删除了
			// len(s)-i 未删除前的总长度，假如i=0,即start =0， 则长度就是len(s)
			if lRemove+rRemove > len(s)-i {
				return
			}
			// 删除左括号
			if lRemove > 0 && s[i] == '(' {
				// 从i开始，继续寻找删除下一个字符
				// 即使字符相同也没关系，因为这一次是再删一个字符，删除'('和删除'(('是不一样的
				dfs(s[:i]+s[i+1:], i, lRemove-1, rRemove)
			}
			// 删除右括号
			if rRemove > 0 && s[i] == ')' {
				dfs(s[:i]+s[i+1:], i, lRemove, rRemove-1)
			}
		}
	}

	//计算总共最少需要删除的左括号和右括号分别的个数
	lRemove, rRemove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lRemove++
		} else if ch == ')' {
			if lRemove > 0 {
				lRemove--
			} else {
				rRemove++
			}
		}
	}
	//回溯遍历
	dfs(s, 0, lRemove, rRemove)
	return ans
}

func TestRemoveInvalidParentheses(t *testing.T) {
	fmt.Println(removeInvalidParentheses("()())()"))
}
