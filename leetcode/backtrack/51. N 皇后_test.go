package backtrack

import (
	"bytes"
	"fmt"
	"testing"
)

/*
*
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例 1：

输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
示例 2：

输入：n = 1
输出：[["Q"]]

题目简介
在 n x n 的棋盘上放置 n 个皇后，使得它们 不能互相攻击（即：同一行、同一列、同一对角线不能同时有两个皇后）。输出所有的解。

解题思路（回溯 Backtracking）
✅ 基本规则：
皇后不能：
在同一行（我们逐行放置，天然不冲突）
在同一列（需要记录列是否被占）
在同一对角线上（主对角线：row - col 相同，副对角线：row + col 相同）

🔁 回溯的核心框架：
从第 0 行开始尝试，在每一列试着放一个皇后；
如果当前位置安全（不冲突），就：
  - 放皇后，
  - 进入下一行递归；

回溯（撤回选择）。

注意：在 n x n 的棋盘上放置 n 个皇后 (每行肯定至少会放一个)

看题解，半写半看
*/
func solveNQueens(n int) [][]string {
	cols := make(map[int]bool)
	diags1 := make(map[int]bool)
	diags2 := make(map[int]bool)
	// 构造初始棋盘
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = bytes.Repeat([]byte("."), n)
	}
	var res [][]string
	var backtrack func(board [][]byte, row int)
	backtrack = func(board [][]byte, row int) {
		if row == n {
			tmp := make([]string, n)
			for i := 0; i < len(board); i++ {
				tmp[i] = string(board[i])
			}
			res = append(res, tmp)
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || diags1[row-col] || diags2[row+col] {
				continue
			}
			board[row][col] = 'Q'
			cols[col], diags1[row-col], diags2[row+col] = true, true, true
			backtrack(board, row+1)
			cols[col], diags1[row-col], diags2[row+col] = false, false, false
			board[row][col] = '.'
		}
	}
	backtrack(board, 0)
	return res
}

func TestSolveNQueens(t *testing.T) {
	fmt.Println(solveNQueens(4))
}
