package leetcode

import (
	"fmt"
	"testing"
)

/**
做了排列和子集之后，尝试做组合
+ exercise/leetcode/hot100/46. 全排列_test.go
+ exercise/leetcode/hot100/78. 子集_test.go

还是不会。。。
chatGPT帮我修正来错误。。

时间
24 ms
击败
6.46%
内存
5.9 MB
击败
81.48%
*/
/**
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。


*/
func combine(n int, k int) [][]int {

	var ans [][]int
	var set []int

	var dfs func(int)
	dfs = func(cur int) {
		if len(set) == k {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		for i := cur; i < n; i++ {
			set = append(set, i+1)
			dfs(i + 1)
			set = set[:len(set)-1]
		}
	}

	dfs(0)
	return ans
}

func TestCombine(t *testing.T) {
	fmt.Println(combine(4, 2))
}
