package dp

import (
	"fmt"
	"testing"
)

/*
*
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。

请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。

假设每一种面额的硬币有无限个。

题目数据保证结果符合 32 位带符号整数。

示例 1：

输入：amount = 5, coins = [1, 2, 5]
输出：4
解释：有四种方式可以凑成总金额：
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
示例 2：

输入：amount = 3, coins = [2]
输出：0
解释：只用面额 2 的硬币不能凑成总金额 3 。
示例 3：

输入：amount = 10, coins = [10]
输出：1

提示：

1 <= coins.length <= 300
1 <= coins[i] <= 5000
coins 中的所有值 互不相同
0 <= amount <= 5000
*/
/**
发现我已经回溯成魔了
*/
/*func change(amount int, coins []int) int {
	count := 0
	var dfs func(start, sum int)
	dfs = func(start, remain int) {
		if start == len(coins) {
			return
		}
		if remain == 0 {
			count++
			return
		}
		for i := start; i < len(coins); i++ {
			if coins[i] <= remain {
				//重复选
				dfs(i, remain-coins[i])
				//选下一个
				//dfs(i+1, remain-coins[i])
			}
			//不选这个
			//dfs(i+1, remain)
		}
	}
	dfs(0, amount)
	return count
}*/

/**
超出时间限制
14 / 30 个通过的测试用例
最后执行的输入
添加到测试用例
amount =
500
coins =
[3,5,7,8,9,10,11]
*/

func change(amount int, coins []int) int {
	count := 0

	return count
}

func TestChange(t *testing.T) {
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(change(3, []int{2}))
	fmt.Println(change(10, []int{10}))
}
