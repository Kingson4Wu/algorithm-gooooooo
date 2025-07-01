package _025

import (
	"fmt"
	"testing"
)

/*
*
目前想法，二维动态规划即可
以左上角为1的情况，判断右方、下方、右下方，是否都是1即可
初始化（长度为1）if matrix[i][j] == 1 dp[i][j] = true
dp[i][j] = dp[i][j] && dp[i][j+1] && dp[i+1][j] && dp[i+1][j+1]
*/
func maximalSquare(matrix [][]byte) int {

	maxVal := 0
	dp := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]bool, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = true
				maxVal = 1
			}
		}
	}

	length := 2

	for {
		change := false
		for i := 0; i <= len(dp)-length; i++ {
			for j := 0; j <= len(dp[i])-length; j++ {
				dp[i][j] = dp[i][j] && dp[i+1][j] && dp[i][j+1] && dp[i+1][j+1]
				if !change && dp[i][j] {
					maxVal = length * length
					change = true
				}
			}
		}
		if !change {
			break
		}
		length++
		if length > len(dp) || length > len(dp[0]) {
			break
		}
	}
	return maxVal
}

func TestMaximalSquare(t *testing.T) {
	/*	fmt.Println(maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}))*/

	fmt.Println(maximalSquare([][]byte{
		{'1', '0', '1', '0'},
		{'1', '0', '1', '1'},
		{'1', '0', '1', '1'},
		{'1', '1', '1', '1'}}))
}

/**
matrix =
[["1","0","1","0"],["1","0","1","1"],["1","0","1","1"],["1","1","1","1"]]

添加到测试用例
输出
9
预期结果
4
*/

/**
执行用时分布
104
ms
击败
5.12%
复杂度分析
消耗内存分布
8.19
MB
击败
95.63%
*/
