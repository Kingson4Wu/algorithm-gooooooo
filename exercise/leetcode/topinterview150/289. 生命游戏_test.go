package topinterview150

import (
	"fmt"
	"testing"
)

/**
个人想法，定义2，3作为临时值，2 表示将要复活，3表示将要死亡
然后遍历每个位置的周围八个位置

执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
3.78
MB
击败
63.29%

*/
/*

根据 百度百科 ， 生命游戏 ，简称为 生命 ，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是 同时 发生的。给你 m x n 网格面板 board 的当前状态，返回下一个状态。

给定当前 board 的状态，更新 board 到下一个状态。

注意 你不需要返回任何东西。
*/
func gameOfLife(board [][]int) {

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

			count := 0
			//左上角
			if i-1 >= 0 && j-1 >= 0 {
				if board[i-1][j-1]%2 == 1 {
					count++

				}
			}
			//上方
			if j-1 >= 0 {
				if board[i][j-1]%2 == 1 {
					count++

				}
			}
			//右上角
			if j-1 >= 0 && i+1 <= len(board)-1 {
				if board[i+1][j-1]%2 == 1 {
					count++

				}
			}
			//右边
			if i+1 <= len(board)-1 {
				if board[i+1][j]%2 == 1 {
					count++

				}
			}
			//右下角
			if j+1 <= len(board[i])-1 && i+1 <= len(board)-1 {
				if board[i+1][j+1]%2 == 1 {
					count++

				}
			}
			//下方
			if j+1 <= len(board[i])-1 {
				if board[i][j+1]%2 == 1 {
					count++

				}
			}
			//左下脚
			if i-1 >= 0 && j+1 <= len(board[i])-1 {
				if board[i-1][j+1]%2 == 1 {
					count++

				}
			}
			//左边
			if i-1 >= 0 {
				if board[i-1][j]%2 == 1 {
					count++
				}
			}

			if board[i][j]%2 == 0 {
				if count == 3 {
					board[i][j] = 2
				}
			}
			if board[i][j]%2 == 1 {
				if count < 2 || count > 3 {
					board[i][j] = 3
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			}
			if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}

/**
执行出错
4 / 23 个通过的测试用例
panic: runtime error: index out of range [1] with length 1
main.gameOfLife({0xc000012210?, 0x7?, 0x8?})
solution.go, line 30
main.__helper__(...)
solution.go, line 91
main.main()
solution.go, line 121
最后执行的输入
添加到测试用例
board =
[[0,0]]
*/

func TestGameOfLife(t *testing.T) {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j])
			fmt.Print(",")
		}
		fmt.Println()
	}
	board = [][]int{{1, 1}, {1, 0}}
	gameOfLife(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j])
			fmt.Print(",")
		}
		fmt.Println()
	}

	board = [][]int{{0, 0}}
	gameOfLife(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j])
			fmt.Print(",")
		}
		fmt.Println()
	}
}
