package topinterview150

import (
	"fmt"
	"testing"
)

/**
写错了一次，了；漏了逻辑分支
双hashmap

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.46
MB
击败
20.06%
复杂度分析
*/
/*
*
给定两个字符串 s 和 t ，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

示例 1:

输入：s = "egg", t = "add"
输出：true
示例 2：

输入：s = "foo", t = "bar"
输出：false
示例 3：

输入：s = "paper", t = "title"
输出：true
*/
func isIsomorphic(s string, t string) bool {

	m := make(map[uint8]uint8)
	rm := make(map[uint8]struct{})
	for i := 0; i < len(s); i++ {
		if ch, ok := m[s[i]]; ok {
			if t[i] != ch {
				return false
			}
		} else {
			// 漏了这步！！不能被重复映射
			if _, ok := rm[t[i]]; ok {
				return false
			}
			m[s[i]] = t[i]
			rm[t[i]] = struct{}{}
		}
	}
	return true
}

/**
解答错误
40 / 47 个通过的测试用例

官方题解
输入
s =
"badc"
t =
"baba"

添加到测试用例
输出
true
预期结果
false
*/

func TestIsIsomorphic(t *testing.T) {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("badc", "baba"))
}
