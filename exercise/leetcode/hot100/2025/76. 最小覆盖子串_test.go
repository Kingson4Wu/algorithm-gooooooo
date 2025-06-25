package _025

import (
	"fmt"
	"testing"
)

/*
*

滑动窗口（增加队列，保存下一次的滑动窗口的left）
hashmap 计算满足的数量 提示判断效率
*/
func minWindow(s string, t string) string {

	var queue []int
	// 保存字符的数量
	charMap := make(map[int32]int)
	//保存字符种类的数量 // 即len(charMap)
	count := 0
	start := 0
	length := 0

	tMap := make(map[int32]int)
	for _, ch := range t {
		if num, ok := tMap[ch]; ok {
			tMap[ch] = num + 1
		} else {
			tMap[ch] = 1
		}
	}

	for i, source := range s {

		if tNum, ok := tMap[source]; ok {
			if num, ok := charMap[source]; ok {
				charMap[source] = num + 1
			} else {
				charMap[source] = 1
			}
			if charMap[source] <= tNum {
				count++
			}
			if len(queue) == 0 {
				start = i
			}
			queue = append(queue, i)

			for count == len(t) {
				size := queue[len(queue)-1] - queue[0] + 1
				if length == 0 || size < length {
					length = size
					start = queue[0]
				}

				prefixChar := int32(s[queue[0]])
				if num, ok := charMap[prefixChar]; ok {
					if num == 1 {
						delete(charMap, prefixChar)
						count--
					} else {
						charMap[prefixChar] = num - 1
						if charMap[prefixChar] < tMap[prefixChar] {
							count--
						}
					}

				}
				queue = queue[1:]
			}
		}

	}

	if length == 0 {
		return ""
	}
	return s[start : start+length]

}

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("aab", "aab"))

	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))

	fmt.Println(minWindow("abbaaavvvaaaaavvva", "aa"))
	fmt.Println(minWindow("abbaaavvvaaaaavvva", "a"))
	fmt.Println(minWindow("abbaaavvvaaaaavvva", "ab"))
	fmt.Println(minWindow("abbaaavvvaaaaavvva", "abvaaaaab"))

}

/**
BANC
a

aa
a
ab
abbaaavvvaa
*/

/**
解答错误
186 / 268 个通过的测试用例

官方题解
输入
s =
"aab"
t =
"aab"

添加到测试用例
输出
"ab"
预期结果
"aab"
*/

/**
执行用时分布
20
ms
击败
69.90%
复杂度分析
消耗内存分布
8.01
MB
击败
5.00%
*/
