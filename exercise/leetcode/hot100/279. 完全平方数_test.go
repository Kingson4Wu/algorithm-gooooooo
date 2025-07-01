package hot100

import (
	"fmt"
	"testing"
)

/**
个人感觉跟零钱兑换是一模一样的，尝试用与之相同的方法做

动态规划
dp[i] 表示最少需要多少个数的平方来表示整数 i
这些数必然落在区间 [1,根号i]
dp[i] = 1 + min(dp[i-j的平方]); j = [1,根号i]

时间
28 ms
击败
59%
内存
6.1 MB
击败
42.66%


官方题解还有数学法，四平方和
*/

/*
*
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

示例 1：

输入：n = 12
输出：3
解释：12 = 4 + 4 + 4
示例 2：

输入：n = 13
输出：2
解释：13 = 4 + 9
*/
func numSquares(n int) int {

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			s := j * j
			if s > i {
				break
			}
			if dp[i] == 0 || dp[i] > dp[i-s]+1 {
				dp[i] = dp[i-s] + 1
			}
		}
	}
	return dp[n]
}

func TestNumSquares(t *testing.T) {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(0))
	fmt.Println(numSquares(1))
	fmt.Println(numSquares(2))
	fmt.Println(numSquares(3))
	fmt.Println(numSquares(4))
	fmt.Println(numSquares(5))
}
