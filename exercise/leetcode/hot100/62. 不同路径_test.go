package hot100

import (
	"fmt"
	"testing"
)

/**
看题解，
1、动态规划
2、组合数学
从左上角到右下角的过程中，我们需要移动 m+n−2 次，其中有 m−1次向下移动，n−1次向右移动。因此路径的总数，就等于从 m+n−2 次移动中选择 m−1 次向下移动的方案数，即组合数
C (m+n−2,m−1)

作者：力扣官方题解
链接：https://leetcode.cn/problems/unique-paths/solutions/514311/bu-tong-lu-jing-by-leetcode-solution-hzjf/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

/**
自己做的，拆分子问题，通过递归可解,但是超时

超出时间限制
41 / 63 个通过的测试用例
最后执行的输入
添加到测试用例
m =
51
n =
9

改成使用dfs回溯 仍然超出时间限制，。。。。。


*/

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？
输入：m = 3, n = 7
输出：28
示例 2：

输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下
示例 3：

输入：m = 7, n = 3
输出：28
示例 4：

输入：m = 3, n = 3
输出：6
*/

// 递归超时
/*func uniquePaths(m int, n int) int {

	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}*/

//回溯也超时
/*func uniquePaths(m int, n int) int {

	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	ans := 0
	var dfs func(int, int)
	dfs = func(mm int, nn int) {

		if mm == m-1 && nn == n-1 {
			ans++
			return
		}
		if mm+1 < m {
			dfs(mm+1, nn)
		}
		if nn+1 < n {
			dfs(mm, nn+1)
		}
	}

	dfs(0, 0)
	return ans
}*/

// 要动态规划。。。
/**
时间
0 ms
击败
100%
内存
1.9 MB
击败
17.87%
*/
func uniquePaths(m int, n int) int {

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func TestUniquePaths(t *testing.T) {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(3, 3))
	fmt.Println(uniquePaths(5, 5))
	fmt.Println(uniquePaths(1, 1))
	fmt.Println(uniquePaths(2, 2))
}
