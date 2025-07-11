package topinterview145

/**

定义三个hash

*/
/**
请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）


注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
空白格用 '.' 表示。

提示：

board.length == 9
board[i].length == 9
board[i][j] 是一位数字（1-9）或者 '.'

*/

/*
	func main() {
		fmt.Println(isValidSudoku([][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))
	}
*/
func isValidSudoku(board [][]byte) bool {

	//9个（3x3）单元格
	m := make([][]map[byte]bool, 3)
	for i := 0; i < 3; i++ {
		m[i] = make([]map[byte]bool, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m[i][j] = map[byte]bool{}
		}
	}

	//9行
	x := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		x[i] = map[byte]bool{}
	}
	//9列
	y := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		y[i] = map[byte]bool{}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}
			//当前行是否已存在改字符
			if x[i][board[i][j]] {
				return false
			}
			x[i][board[i][j]] = true
			//当前列是否已存在改字符
			if y[j][board[i][j]] {
				return false
			}
			y[j][board[i][j]] = true
			if m[i/3][j/3][board[i][j]] {
				return false
			}
			m[i/3][j/3][board[i][j]] = true
		}
	}
	return true
}
