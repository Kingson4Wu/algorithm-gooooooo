package hot100

import (
	"fmt"
	"testing"
)

/**
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false
*/

/**
看了题解之后回忆

递归、回溯、动态规划、剪枝

本质上就是穷举
回溯总是很难理解（回溯恢复状态，要记住，整个流程都是串行，没有并发问题）
动态规划定义的公式也很难想到

回溯
思路与算法

设函数 check(i,j,k) 表示判断以网格的 (i,j) 位置出发，能否搜索到单词 word[k..]，其中 word[k..]表示字符串 word从第 k 个字符开始的后缀子串。
如果能搜索到，则返回 true，反之返回 false。
1. 如果 board[i][j]≠s[k]，当前字符不匹配，直接返回 false\texttt{false}false。
2. 如果当前已经访问到字符串的末尾，且对应字符依然匹配，此时直接返回 true。
3. 否则，遍历当前位置的所有相邻位置。如果从某个相邻位置出发，能够搜索到子串 word[k+1..]，则返回 true，否则返回 false。（四个方向）

对每一个位置 (i,j) 都调用函数 check(i,j,0) 进行检查：只要有一处返回 true，就说明网格中能够找到相应的单词，否则说明不能找到

为了防止重复遍历相同的位置，需要额外维护一个与 board 等大的 visited 数组，用于标识每个位置是否被访问过。每次遍历相邻位置时，需要跳过已经被访问的位置。
（记录每一个位置作为起点时，走过的位置，避免重复）


*/

func exist(board [][]byte, word string) bool {

	return false
}

func TestExist(t *testing.T) {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
}
