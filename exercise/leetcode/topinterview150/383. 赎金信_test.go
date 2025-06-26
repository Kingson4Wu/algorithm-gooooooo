package topinterview150

import (
	"fmt"
	"testing"
)

/**
执行用时分布
9
ms
击败
47.56%
复杂度分析
消耗内存分布
5.60
MB
击败
49.19%
*/
/*
*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。

示例 1：

输入：ransomNote = "a", magazine = "b"
输出：false
示例 2：

输入：ransomNote = "aa", magazine = "ab"
输出：false
示例 3：

输入：ransomNote = "aa", magazine = "aab"
输出：true
*/
func canConstruct(ransomNote string, magazine string) bool {

	m := make(map[rune]int)
	for _, c := range magazine {
		m[c]++
	}
	for _, c := range ransomNote {
		if num, ok := m[c]; ok {
			if num == 0 {
				return false
			}
			m[c]--
		} else {
			return false
		}
	}

	return true
}

func TestCanConstruct(t *testing.T) {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
