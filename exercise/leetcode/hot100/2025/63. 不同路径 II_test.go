package _025

import (
	"fmt"
	"testing"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	//第一行和第一列都只能从一个方向走过来，所以都为 1
	//注意障碍物
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	} else {
		return 0
	}
	pass := false
	if dp[0][0] == 1 {
		for i := 1; i < len(dp); i++ {
			if obstacleGrid[i][0] == 0 && dp[i-1][0] == 1 {
				dp[i][0] = 1
				pass = true
			} else {
				break
			}
		}
		for j := 1; j < len(dp[0]); j++ {
			if obstacleGrid[0][j] == 0 && dp[0][j-1] == 1 {
				dp[0][j] = 1
				pass = true
			} else {
				break
			}
		}
	}
	if !pass {
		return dp[m-1][n-1]
	}

	// 怎么提前剪枝？？！
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]
}

/**
解答错误
25 / 42 个通过的测试用例

官方题解
输入
obstacleGrid =
[[1,0]]

添加到测试用例
输出
1
预期结果
0

解答错误
41 / 42 个通过的测试用例

官方题解
输入
obstacleGrid =
[[0]]

添加到测试用例
输出
0
预期结果
1

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.15
MB
击败
70.72%
*/

func TestUniquePathsWithObstacles(t *testing.T) {
	fmt.Println(uniquePathsWithObstacles([][]int{{0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{1, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}
