package hot100

import (
	"fmt"
	"testing"
)

/**
自己做的，其实就类似于栈的思想
*/

/*
*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false

提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/
func isValid(s string) bool {

	var leftArr []rune

	for _, c := range s {

		if c == '(' || c == '{' || c == '[' {
			leftArr = append(leftArr, c)
		} else {
			if len(leftArr) == 0 {
				return false
			}
			lc := leftArr[len(leftArr)-1]
			if lc == '(' && c != ')' {
				return false
			}
			if lc == '{' && c != '}' {
				return false
			}
			if lc == '[' && c != ']' {
				return false
			}
			leftArr = leftArr[0 : len(leftArr)-1]
		}
	}

	return len(leftArr) == 0
}

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("((({}))){}[]{{[({}{}(({})))]}}"))
	fmt.Println(isValid("((({}))){}[]{{[({}{}(({)))]}}}"))
}
