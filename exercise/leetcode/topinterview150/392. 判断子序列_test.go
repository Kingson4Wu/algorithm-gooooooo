package topinterview150

import (
	"fmt"
	"testing"
)

/*
*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

进阶：

如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
*/
func isSubsequence(s string, t string) bool {

	index := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[index] {
			index++
		} else {
			continue
		}
		if index == len(s) {
			return true
		}
	}
	return false
}

/**
panic: runtime error: index out of range [0] with length 0
main.isSubsequence(...)
solution.go, line 5
main.__helper__(...)
solution.go, line 20
main.main()
solution.go, line 61
最后执行的输入
添加到测试用例
s =
""
t =
"ahbgdc"
*/

func TestIsSubsequence(t *testing.T) {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
