package topinterview145

import (
	"fmt"
	"testing"
)

/**
给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。

对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。

示例 1：

输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
输出：4
解释：最长递增路径为 [1, 2, 6, 9]。
示例 2：

输入：matrix = [[3,4,5],[3,2,6],[2,2,1]]
输出：4
解释：最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
示例 3：

输入：matrix = [[1]]
输出：1

提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
0 <= matrix[i][j] <= 231 - 1
*/

/*
*
感觉是图的深度优先搜索
结束递归是上下左右的数都大于等于当前，直接返回1
记录当前结点的结果，并作为visited标记

执行用时分布
13
ms
击败
10.66%
复杂度分析
消耗内存分布
8.52
MB
击败
23.77%
*/
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	pathCount := make([][]int, m)
	for i := 0; i < m; i++ {
		pathCount[i] = make([]int, n)
	}
	//是否还能继续走，有更大的值
	check := func(i, j int) bool {
		if i-1 >= 0 && matrix[i-1][j] > matrix[i][j] {
			return true
		}
		if i+1 < m && matrix[i+1][j] > matrix[i][j] {
			return true
		}
		if j-1 >= 0 && matrix[i][j-1] > matrix[i][j] {
			return true
		}
		if j+1 < n && matrix[i][j+1] > matrix[i][j] {
			return true
		}
		return false
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var ans int
	var dfs func(i, j, preNum int) int
	dfs = func(i, j, preNum int) int {
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		if preNum >= matrix[i][j] {
			return 0
		}
		//相当于visited标记
		if pathCount[i][j] > 0 {
			ans = max(ans, 1)
			return pathCount[i][j]
		}
		if !check(i, j) {
			pathCount[i][j] = 1
			ans = max(ans, 1)
			return pathCount[i][j]
		}
		//上下左右
		pathCount[i][j] = 1 + max(
			max(dfs(i-1, j, matrix[i][j]), dfs(i+1, j, matrix[i][j])),
			max(dfs(i, j-1, matrix[i][j]), dfs(i, j+1, matrix[i][j])))
		ans = max(ans, pathCount[i][j])
		return pathCount[i][j]
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			dfs(i, j, -1)
		}
	}
	return ans
}

func TestLongestIncreasingPath(t *testing.T) {
	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
	fmt.Println(longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}))
	fmt.Println(longestIncreasingPath([][]int{{1}}))
}
