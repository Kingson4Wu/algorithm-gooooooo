package topinterview150

import (
	"fmt"
	"math"
	"testing"
)

/*
*
给定一个三角形 triangle ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。

示例 1：

输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：

	  2
	 3 4
	6 5 7

4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
示例 2：

输入：triangle = [[-10]]
输出：-10

提示：

1 <= triangle.length <= 200
triangle[0].length == 1
triangle[i].length == triangle[i - 1].length + 1
-104 <= triangle[i][j] <= 104

进阶：

你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
*/
/**
感觉用回溯
超出时间限制！！！！
*/
/*func minimumTotal(triangle [][]int) int {

	minSum := math.MaxInt32
	sum := triangle[0][0]
	var dfs func(level, i, j int)
	dfs = func(level, i, j int) {
		if level == len(triangle) {
			if sum < minSum {
				minSum = sum
			}
			return
		}
		//选左
		sum += triangle[i+1][j]
		dfs(level+1, i+1, j)
		sum -= triangle[i+1][j]
		//选右
		sum += triangle[i+1][j+1]
		dfs(level+1, i+1, j+1)
		sum -= triangle[i+1][j+1]
	}
	dfs(1, 0, 0)
	return minSum
}*/

/**
超出时间限制
42 / 45 个通过的测试用例
最后执行的输入
添加到测试用例
triangle =
[[0],[0,0],[0,0,0],[0,0,0,0],[0,0,0,0,0],[0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[0,0,0,0,0,
*/

// 动态规划 不是很难想
/**

最左边加上上一层的第一个 (0)
最右边加上上一层的最后一个（j-1）
中间的选 0、j-1中最小的一个

dp := make([]int, level)
dp[0] = triangle[0][0]

if j == 0 {
	dp[j] = triangle[i][j] + dp[j]
} else if j == i {
	dp[j] = triangle[i][j] + dp[j-1]
} else {
	dp[j] = triangle[i][j] + min(dp[j-1], dp[j])
}

-----

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
4.70
MB
击败
72.55%

*/
func minimumTotal(triangle [][]int) int {
	level := len(triangle)
	if level == 1 {
		return triangle[0][0]
	}

	minSum := math.MaxInt32
	dp := make([]int, level)
	dp[0] = triangle[0][0]

	for i := 1; i < level; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 {
				dp[j] = triangle[i][j] + dp[j]
			} else if j == i {
				dp[j] = triangle[i][j] + dp[j-1]
			} else {
				dp[j] = triangle[i][j] + min(dp[j-1], dp[j])
			}
			if i == level-1 {
				if dp[j] < minSum {
					minSum = dp[j]
				}
			}
		}
	}
	return minSum
}

func TestMinimumTotal(t *testing.T) {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
	fmt.Println(minimumTotal([][]int{{-10}}))
}
