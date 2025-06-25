package hot100

import (
	"fmt"
	"testing"
)

/*
*
leetcode/dp/64. 最小路径和.go

给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

执行用时分布
1
ms
击败
17.05%
复杂度分析
消耗内存分布
6.54
MB
击败
79.56%

// 动态规划
// 遍历每一行，参考上和左算出最小值即可
// 还可以优化空间复杂度为O（1），在原数组上面操作
*/
func minPathSum(grid [][]int) int {

	// 我这次似乎终于直接想到节约空间的最优解
	n := len(grid)
	m := len(grid[0])
	dp := make([]int, m)

	//初始化dp数据
	dp[0] = grid[0][0]
	for j := 1; j < m; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	for i := 1; i < n; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < m; j++ {
			if dp[j-1] < dp[j] {
				dp[j] = dp[j-1] + grid[i][j]
			} else {
				dp[j] += grid[i][j]
			}
		}

	}
	return dp[m-1]
}

func TestMinPathSum(t *testing.T) {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
