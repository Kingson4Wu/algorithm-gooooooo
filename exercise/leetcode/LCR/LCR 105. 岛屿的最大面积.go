package LCR

/*
*
自己按照岛屿个数的思路，举一反三做出来了

执行用时分布
15
ms
击败
25.76%
复杂度分析
消耗内存分布
5.70
MB
击败
98.48
*/
func maxAreaOfIsland(grid [][]int) int {

	count := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		count++
		grid[i][j] = 0
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	maxCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			count = 0
			dfs(i, j)
			if count > maxCount {
				maxCount = count
			}
		}
	}
	return maxCount
}
