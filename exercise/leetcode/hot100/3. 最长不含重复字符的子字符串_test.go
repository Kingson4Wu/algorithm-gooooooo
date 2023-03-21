package hot100

import (
	"fmt"
	"testing"
)

/** 双指针（即滑动窗口）+ hashmap */
/** 我的实现不是双指针，但是跟双指针的殊途同归，使用current变量 */

/**
s =
" "
添加到测试用例
输出
0
预期结果
1

Go
时间
8 ms
击败
69.47%
内存
2.6 MB
击败
73.49%

*/

/** 我的实现不是更好吗？ */
func lengthOfLongestSubstring(s string) int {

	m := make(map[byte]int)
	max := 0
	current := 0

	for i := 0; i < len(s); i++ {

		if _, ok := m[s[i]]; !ok {
			current++
			m[s[i]] = 0
		} else {
			if current > max {
				max = current
			}
			for j := i - current; j < i; j++ {
				if s[j] == s[i] {
					break
				}
				delete(m, s[j])
				current--
			}
		}
	}
	if current > max {
		max = current
	}

	return max
}

func TestLengthOfLongestSubstring(t *testing.T) {

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(" "))

}
