package hot100

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
*
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：



输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"

提示：

1 <= s.length <= 30
s 由小写英文字母、数字和方括号 '[]' 组成
s 保证是一个 有效 的输入。
s 中所有整数的取值范围为 [1, 300]
*/

/**
自己思考后的想法：
1、使用栈的思想遍历时，存储左括号的下标
*/

/**
方法一：栈操作
方法二：递归

一看就会，写起来就废系列

方法一：栈操作

func decodeString(s string) string {
    stk := []string{}
    ptr := 0
    for ptr < len(s) {
        cur := s[ptr]
        if cur >= '0' && cur <= '9' {
            digits := getDigits(s, &ptr)
            stk = append(stk, digits)
        } else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
            stk = append(stk, string(cur))
            ptr++
        } else {
            ptr++
            sub := []string{}
            for stk[len(stk)-1] != "[" {
                sub = append(sub, stk[len(stk)-1])
                stk = stk[:len(stk)-1]
            }
            for i := 0; i < len(sub)/2; i++ {
                sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
            }
            stk = stk[:len(stk)-1]
            repTime, _ := strconv.Atoi(stk[len(stk)-1])
            stk = stk[:len(stk)-1]
            t := strings.Repeat(getString(sub), repTime)
            stk = append(stk, t)
        }
    }
    return getString(stk)
}

func getDigits(s string, ptr *int) string {
    ret := ""
    for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
        ret += string(s[*ptr])
    }
    return ret
}

func getString(v []string) string {
    ret := ""
    for _, s := range v {
        ret += s
    }
    return ret
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/decode-string/solutions/264391/zi-fu-chuan-jie-ma-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

方法二：递归

var (
    src string
    ptr int
)

func decodeString(s string) string {
    src = s
    ptr = 0
    return getString()
}

func getString() string {
    if ptr == len(src) || src[ptr] == ']' {
        return ""
    }
    cur := src[ptr]
    repTime := 1
    ret := ""
    if cur >= '0' && cur <= '9' {
        repTime = getDigits()
        ptr++
        str := getString()
        ptr++
        ret = strings.Repeat(str, repTime)
    } else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {
        ret = string(cur)
        ptr++
    }
    return ret + getString()
}

func getDigits() int {
    ret := 0
    for ; src[ptr] >= '0' && src[ptr] <= '9'; ptr++ {
        ret = ret * 10 + int(src[ptr] - '0')
    }
    return ret
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/decode-string/solutions/264391/zi-fu-chuan-jie-ma-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

/*
*
自己想的，用栈实现
双栈法
栈保存重复的次数(数字前缀)
对应的栈下标的辅助栈保存将要重复的字符串
遇到'['入栈，遇到']'出栈， 出栈 追加到 栈顶 次数对应的辅助栈 的前缀字符串，如果栈为空，则追加到最终的结果字符串
*/
func decodeString(s string) string {

	//3[a2[c]]
	//2[abc]3[cd]ef

	var stack []int
	var prestr []string
	var result string

	var nums []rune
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			if len(stack) == 0 {
				result += string(ch)
			} else {
				prestr[len(prestr)-1] += string(ch)
			}
		} else if ch >= '0' && ch <= '9' {
			nums = append(nums, ch)
		} else if ch == '[' {
			count, _ := strconv.Atoi(string(nums))
			nums = nums[:0]
			stack = append(stack, count)
			prestr = append(prestr, "")
		} else if ch == ']' {
			str := strings.Repeat(prestr[len(prestr)-1], stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			prestr = prestr[:len(prestr)-1]
			if len(stack) == 0 {
				result += str
			} else {
				prestr[len(prestr)-1] += str
			}
		}
	}
	return result
}

/**
解答错误
28 / 34 个通过的测试用例

官方题解
输入
s =
"100[leetcode]"

添加到测试用例
输出
"leetcode"

else if ch >= '1' && ch <= '9' {
改成
else if ch >= '0' && ch <= '9' {

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.92
MB
击败
41.74%
*/

func TestDecodeString(t *testing.T) {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("abc3[cd]xyz"))
}
