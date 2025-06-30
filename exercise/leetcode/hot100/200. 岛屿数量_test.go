package hot100

import (
	"fmt"
	"testing"
)

/*
*

给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

示例 1：

输入：grid = [

	["1","1","1","1","0"],
	["1","1","0","1","0"],
	["1","1","0","0","0"],
	["0","0","0","0","0"]

]
输出：1
示例 2：

输入：grid = [

	["1","1","0","0","0"],
	["1","1","0","0","0"],
	["0","0","1","0","0"],
	["0","0","0","1","1"]

]
输出：3

提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
*/
/**
看过题解两天后回忆写的
将二维网格看成一个无向图，竖直或水平相邻的 1 之间有边相连
无向图深度优先搜索
将1变成0的过程相当于标记visited
遍历所有结点，是1的统计加1，深度搜索将周围的1全变成0，最后统计出来的count即岛屿数量
*/
func numIslands(grid [][]byte) int {

	count := 0
	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if i < 0 || i >= len(grid) {
			return
		}
		if j < 0 || j >= len(grid[i]) {
			return
		}
		if grid[i][j] == '1' {

			grid[i][j] = '0'
			// 向左
			dfs(grid, i, j-1)
			//向右
			dfs(grid, i, j+1)
			//向上
			dfs(grid, i-1, j)
			//向下
			dfs(grid, i+1, j)
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

/**
执行用时分布
3
ms
击败
75.32%
复杂度分析
消耗内存分布
5.47
MB
击败
83.92%

*/

func TestNumIslands(t *testing.T) {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
	fmt.Println(numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}))

}
