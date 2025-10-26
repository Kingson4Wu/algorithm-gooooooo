package LCR

import (
	"fmt"
	"testing"
)

// 无向图的深度优先搜索
/**
自己做的
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
5.25
MB
击败
50.98%

*/
func wordPuzzle(grid [][]byte, target string) bool {

	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}
	result := false
	var dfs func(i, j int, target string)
	dfs = func(i, j int, target string) {
		if result {
			return
		}
		if target == "" {
			result = true
			return
		}
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
			return
		}
		if visited[i][j] {
			return
		}

		if target[0] == grid[i][j] {
			visited[i][j] = true
			dfs(i-1, j, target[1:])
			dfs(i+1, j, target[1:])
			dfs(i, j-1, target[1:])
			dfs(i, j+1, target[1:])
			visited[i][j] = false
		} else {
			return
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if result {
				return true
			}
			dfs(i, j, target)
		}
	}

	return result
}

func TestWordPuzzle(t *testing.T) {
	fmt.Println(wordPuzzle([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
}
