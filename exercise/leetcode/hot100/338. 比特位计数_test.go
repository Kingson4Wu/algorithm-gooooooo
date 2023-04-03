package hot100

import (
	"fmt"
	"testing"
)

/**
看题解做的

时间
4 ms
击败
82.78%
内存
4.3 MB
击败
81.39%

方法三：动态规划——最低有效位


*/

/*
*
给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。

示例 1：

输入：n = 2
输出：[0,1,1]
解释：
0 --> 0
1 --> 1
2 --> 10
示例 2：

输入：n = 5
输出：[0,1,1,2,1,2]
解释：
0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101

进阶：

很容易就能实现时间复杂度为 O(n log n) 的解决方案，你可以在线性时间复杂度 O(n) 内用一趟扫描解决此问题吗？
你能不使用任何内置函数解决此问题吗？（如，C++ 中的 __builtin_popcount ）
*/
func countBits(n int) []int {

	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + i&1
	}

	return dp
}

func TestCountBits(t *testing.T) {
	fmt.Println(countBits(10))
}
