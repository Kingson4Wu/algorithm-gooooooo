package topinterview150

/*
*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

进阶：

一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个仅使用常量空间的解决方案吗？
*/
/**
用m+n两个数组记录要置0的行和列
再遍历一次，将符合的置0

其他优化：两个变量记录第一行和第一列是否原来有0，然后利用第一行和第一列来记录
1个变量，调整遍历顺序即可，太骚
*/
func setZeroes(matrix [][]int) {

	m := len(matrix)
	n := len(matrix[0])
	rows := make([]bool, m)
	cols := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

/**
执行用时分布
0
ms
击败
100.00%
复杂度分析
消耗内存分布
7.31
MB
击败
50.79%

*/
