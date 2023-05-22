package hot100

import (
	"fmt"
	"testing"
)

/**
方法一：动态规划
难理解
时间复杂度： O(n)
空间复杂度： O(n)
方法二：栈
需要想清楚
时间复杂度： O(n)
空间复杂度： O(n)
方法三：不需要额外的空间（算左右数量，反过来再算一遍）

在此方法中，我们利用两个计数器 left 和 right 。首先，我们从左到右遍历字符串，对于遇到的每个 ‘(’，我们增加 left 计数器，对于遇到的每个 ‘)’ ，我们增加 right 计数器。每当 left 计数器与 right 计数器相等时，我们计算当前有效字符串的长度，并且记录目前为止找到的最长子字符串。当 right 计数器比 left 计数器大时，我们将 left 和 right 计数器同时变回 0。

这样的做法贪心地考虑了以当前字符下标结尾的有效括号长度，每次当右括号数量多于左括号数量的时候之前的字符我们都扔掉不再考虑，重新从下一个字符开始计算，但这样会漏掉一种情况，就是遍历的时候左括号的数量始终大于右括号的数量，即 (() ，这种时候最长有效括号是求不出来的。

解决的方法也很简单，我们只需要从右往左遍历用类似的方法计算即可，只是这个时候判断条件反了过来：

当 left 计数器比 right 计数器大时，我们将 left 和 right 计数器同时变回 0
当 left 计数器与 right 计数器相等时，我们计算当前有效字符串的长度，并且记录目前为止找到的最长子字符串

作者：力扣官方题解
链接：https://leetcode.cn/problems/longest-valid-parentheses/solutions/314683/zui-chang-you-xiao-gua-hao-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/**
好像不困难？错觉？

改了两次没好。。

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
/*func longestValidParentheses(s string) int {

	leftNum := 0
	currentLength := 0
	preLength := 0
	longest := 0
	for _, c := range s {
		if c == '(' {
			leftNum++
			continue
		}
		if c == ')' {
			if leftNum > 0 {
				currentLength++
				if currentLength > longest {
					longest = currentLength
				}
				leftNum--
				if leftNum == 0 {
					if preLength > 0 {
						totalLength := currentLength + preLength
						if totalLength > longest {
							longest = totalLength
						}
						preLength = totalLength
						currentLength = 0
					} else {
						preLength = currentLength
						currentLength = 0
					}
				}
			} else {
				leftNum = 0
				currentLength = 0
				preLength = 0
			}
		}
	}

	return longest * 2
}*/

/*
*
时间
0 ms
击败
100%
内存
2.1 MB
击败
98.5%

看了题解用第3种方法做，简单
*/
func longestValidParentheses(s string) int {

	left := 0
	right := 0
	maxLength := 0

	/** 从左到右 */
	for i := 0; i < len(s); i++ {
		if right > left {
			right = 0
			left = 0
		}

		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			right++
		}
		if right == left {
			if right*2 > maxLength {
				maxLength = right * 2
			}
		}
	}

	/** 从右到左 */
	left = 0
	right = 0
	for i := len(s) - 1; i >= 0; i-- {

		if left > right {
			right = 0
			left = 0
		}

		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			right++
		}
		if right == left {
			if right*2 > maxLength {
				maxLength = right * 2
			}
		}
	}

	return maxLength
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
	fmt.Println(longestValidParentheses("(()(((()"))

	/**
	2
	4
	4
	0
	6
	8
	10
	2
	*/
}
