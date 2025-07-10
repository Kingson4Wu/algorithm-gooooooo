package _025

import (
	"fmt"
	"testing"
)

/*
*
二维网格搜索 + 回溯 + DFS + 访问标记

自己做的

执行用时分布
125
ms
击败
53.82%
复杂度分析
消耗内存分布
3.95
MB
击败
32.07%
*/
func exist(board [][]byte, word string) bool {

	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, n)
	}

	flag := false
	var dfs func(i, j, index int)
	dfs = func(i, j, index int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if flag {
			return
		}
		if visited[i][j] {
			return
		}
		if board[i][j] != word[index] {
			return
		}

		if index == len(word)-1 {
			flag = true
			return
		}
		//上下左右
		visited[i][j] = true
		defer func() {
			visited[i][j] = false
		}()
		dfs(i-1, j, index+1)
		dfs(i+1, j, index+1)
		dfs(i, j-1, index+1)
		dfs(i, j+1, index+1)

	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(i, j, 0)
			if flag {
				return true
			}
		}
	}
	return flag
}

func TestExist(t *testing.T) {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
	//无语，细节上犯了两个低级错误
	fmt.Println(exist([][]byte{{'a', 'b'}}, "ba"))
	fmt.Println(exist([][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}, "AAB"))
}

/**
true
false
true
true
*/
