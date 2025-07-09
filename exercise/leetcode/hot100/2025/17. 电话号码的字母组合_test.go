package _025

import (
	"fmt"
	"testing"
)

/*
*
自己用回溯这次是写出来了，但感觉写得不好
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.83
MB
击败
96.28%
*/
func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	var phoneMap = map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var results []string
	var dfs func(start int, str string)
	dfs = func(start int, str string) {
		if start == len(digits) {
			results = append(results, str)
			return
		}
		//这里不是枚举位置，而是枚举当前一个位置的合法字符选项。
		for _, ch := range phoneMap[rune(digits[start])] {
			dfs(start+1, str+string(ch))
		}
	}
	dfs(0, "")
	return results
}

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations(""))
}
