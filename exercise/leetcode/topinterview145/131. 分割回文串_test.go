package topinterview145

import (
	"fmt"
	"testing"
)

/**
给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：

输入：s = "a"
输出：[["a"]]

提示：

1 <= s.length <= 16
s 仅由小写英文字母组成
*/
/**
个人想法（官网方法一和我一样）
动态规划+回溯
自己做的
*/
func partition(s string) [][]string {

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for length := 2; length <= len(s); length++ {
		for i := 0; i < len(s)-length+1; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				if length == 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
		}
	}

	var ans [][]string
	var path []string
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		step := 0
		for start+step < len(s) {
			end := start + step
			if dp[start][end] {
				path = append(path, s[start:end+1])
				dfs(end + 1)
				path = path[:len(path)-1]
			}
			step++
		}
	}
	dfs(0)
	return ans
}

/**
解答错误
24 / 32 个通过的测试用例

官方题解
输入
s =
"efe"

添加到测试用例
输出
[["e","f","e"]]
预期结果
[["e","f","e"],["efe"]]
*/
/**
执行用时分布
17
ms
击败
49.35%
复杂度分析
消耗内存分布
23.66
MB
击败
73.65%

*/

func TestPartition(t *testing.T) {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
	fmt.Println(partition("efe"))
}
