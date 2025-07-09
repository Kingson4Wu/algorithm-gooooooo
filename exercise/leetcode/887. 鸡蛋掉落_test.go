package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/**
给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。

已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。

每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。

请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？


示例 1：

输入：k = 1, n = 2
输出：2
解释：
鸡蛋从 1 楼掉落。如果它碎了，肯定能得出 f = 0 。
否则，鸡蛋从 2 楼掉落。如果它碎了，肯定能得出 f = 1 。
如果它没碎，那么肯定能得出 f = 2 。
因此，在最坏的情况下我们需要移动 2 次以确定 f 是多少。
示例 2：

输入：k = 2, n = 6
输出：3
示例 3：

输入：k = 3, n = 14
输出：4


提示：

1 <= k <= 100
1 <= n <= 104

https://leetcode.cn/problems/super-egg-drop/description/
*/

/*
*
不是二分查找？？
直接看题解
方法一：动态规划 + 二分查找
---
看完题解后写的

方法三：数学法
用 f(t, k) 表示：在最多扔 t 次、拥有 k 个鸡蛋的情况下，最多能测多少层楼。
f(t, k) = 1 + f(t-1, k-1) + f(t-1, k)
它的含义是：
  - 当前扔一层楼（那 1 就是“当前这层”）
  - 碎了 → 剩 t-1 次机会、k-1 个鸡蛋 → 能搞定 f(t-1, k-1) 层 （只能在这层以下的楼层继续测试）
  - 没碎 → 剩 t-1 次机会、k 个鸡蛋 → 能搞定 f(t-1, k) 层 （可以往上继续试）

反过来问：如果我最多只能扔 t 次鸡蛋，手上有 k 个鸡蛋，最多能搞定多少层楼？（就是最大多少层，我一定能找出哪一层开始鸡蛋会碎）
边界条件为：当 t≥1 的时候 f(t,1)=t，当 k≥1 时，f(1,k)=1。

t<=n, 就算你鸡蛋再多，最多是n次就行

	//这句别漏了 ！！！
	if n == 1 {
		return 1
	}
*/
func superEggDrop(k int, n int) int {
	ans := math.MaxInt32
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][1] = 1
	}
	for j := 1; j <= k; j++ {
		dp[1][j] = 1
	}
	//这句别漏了 ！！！
	if n == 1 {
		return 1
	}
	i := 2
	for i <= n {
		for j := 1; j <= k; j++ {
			dp[i][j] = 1 + dp[i-1][j-1] + dp[i-1][j]
			if dp[i][j] >= n && i < ans {
				ans = i
			}
		}
		i++
	}
	return ans
}

/*
*
输入
k =
1
n =
1

添加到测试用例
输出
2147483647
预期结果
1
*/
/**
执行用时分布
34
ms
击败
29.73%
复杂度分析
消耗内存分布
14.18
MB
击败
5.40%
*/
func TestSuperEggDrop(t *testing.T) {
	fmt.Println(superEggDrop(1, 1))
	fmt.Println(superEggDrop(1, 2))
	fmt.Println(superEggDrop(2, 6))
	fmt.Println(superEggDrop(3, 14))
}
