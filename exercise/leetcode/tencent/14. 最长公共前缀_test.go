package tencent

import (
	"fmt"
	"testing"
)

/**
似乎确实不难

时间
0 ms
击败
100%
内存
2.2 MB
击败
26.35%

以前写的是什么垃圾。。
nowcoder/BM84 最长公共前缀.go
*/

/*
*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/
func longestCommonPrefix(strs []string) string {

	if len(strs) == 1 {
		return strs[0]
	}

	index := 0
	keep := true

	for keep {
		for i := 0; i < len(strs)-1; i++ {
			if index >= len(strs[i]) || index >= len(strs[i+1]) || strs[i][index] != strs[i+1][index] {
				keep = false
				index--
				break
			}
		}
		if keep {
			index++
		}
	}

	if index >= 0 {
		return strs[0][:index+1]
	}

	return ""
}

func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))

	fmt.Println(longestCommonPrefix([]string{"d", "defd"}))
}
