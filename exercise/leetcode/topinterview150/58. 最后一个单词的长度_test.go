package topinterview150

import (
	"fmt"
	"testing"
)

/*
*
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。

示例 1：

输入：s = "Hello World"
输出：5
解释：最后一个单词是“World”，长度为 5。
示例 2：

输入：s = "   fly me   to   the moon  "
输出：4
解释：最后一个单词是“moon”，长度为 4。
示例 3：

输入：s = "luffy is still joyboy"
输出：6
解释：最后一个单词是长度为 6 的“joyboy”。
*/
func lengthOfLastWord(s string) int {

	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count > 0 {
				return count
			}
			continue
		}
		count++
	}
	return 0
}

/**
解答错误
51 / 60 个通过的测试用例

官方题解
输入
s =
"a"

添加到测试用例
输出
0
预期结果
1
*/

func TestLengthOfLastWord(t *testing.T) {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}
