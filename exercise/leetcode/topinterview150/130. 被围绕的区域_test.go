package topinterview150

/*
*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' 组成，捕获 所有 被围绕的区域：

连接：一个单元格与水平或垂直方向上相邻的单元格连接。
区域：连接所有 'O' 的单元格来形成一个区域。
围绕：如果您可以用 'X' 单元格 连接这个区域，并且区域中没有任何单元格位于 board 边缘，则该区域被 'X' 单元格围绕。
通过 原地 将输入矩阵中的所有 'O' 替换为 'X' 来 捕获被围绕的区域。你不需要返回任何值。

示例 1：

输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]

输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]

解释：

在上图中，底部的区域没有被捕获，因为它在 board 的边缘并且不能被围绕。

示例 2：

输入：board = [["X"]]

输出：[["X"]]

提示：

m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] 为 'X' 或 'O'
*/
/**
自己想的
无向图遍历，只将边缘结点作为开始遍历（因为这张才没被包围）
向四边深度搜索，遇到O就标记
visited m*n，
再遍历一遍，将为O且未被标记的设置为X
*/
func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var seek func(i, j int)
	seek = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		if board[i][j] != 'O' {
			return
		}
		//上下左右
		seek(i-1, j)
		seek(i+1, j)
		seek(i, j-1)
		seek(i, j+1)
	}
	//上边缘
	for j := 0; j < n; j++ {
		seek(0, j)
	}
	//右边缘
	for i := 0; i < m; i++ {
		seek(i, n-1)
	}
	//下边缘
	for j := 0; j < n; j++ {
		seek(m-1, j)
	}
	//左边缘
	for i := 0; i < m; i++ {
		seek(i, 0)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

/**
执行用时分布
3
ms
击败
24.53%
复杂度分析
消耗内存分布
7.93
MB
击败
94.13%

*/
