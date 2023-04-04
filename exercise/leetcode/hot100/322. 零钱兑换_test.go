package hot100

import (
	"fmt"
	"testing"
)

/**
又是看了题解才会做

我们采用自下而上的方式进行思考。仍定义 F(i) 为组成金额 i 所需最少的硬币数量，假设在计算 F(i) 之前，我们已经计算出 F(0)−F(i−1) 的答案。 则 F(i) 对应的转移方程应为

F(i)=min(j=0…n−1F(i−cj)+1F(i)) + 1

  代表的是第 j 枚硬币的面值，即我们枚举最后一枚硬币面额是 cj
 ，那么需要从 i−cj
  这个金额的状态 F(i−cj)
'
 ) 转移过来，再算上枚举的这枚硬币数量 1 的贡献，由于要硬币数量最少，所以 F(i) 为前面能转移过来的状态的最小值加上枚举的硬币数量 1 。

作者：力扣官方题解
链接：https://leetcode.cn/problems/coin-change/solutions/132979/322-ling-qian-dui-huan-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。



示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0
*/

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				if dp[i] == -1 || dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	return dp[amount]
}

func TestCoinChange(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
}
