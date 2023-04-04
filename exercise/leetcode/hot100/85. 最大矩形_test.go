package hot100

import (
	"fmt"
	"testing"
)

/**
个人想法， 做过最大正方形，这道题应该可以按照同样的套路
动态规划，值是该点是1时作为右下角，最大的长和宽
注意不符合条件的要设置成[]int{0, 0}
*/

/*
*
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。
示例 2：

输入：matrix = []
输出：0
示例 3：

输入：matrix = [["0"]]
输出：0
示例 4：

输入：matrix = [["1"]]
输出：1
示例 5：

输入：matrix = [["0","0"]]
输出：0
*/
func maximalRectangle(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}

	maxArea := 0
	dp := make([][][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([][]int, len(matrix[i]))
	}

	/** 初始化 */
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = []int{1, 1}
			maxArea = 1
		} else {
			dp[i][0] = []int{0, 0}
		}

	}
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = []int{1, 1}
			maxArea = 1
		} else {
			dp[0][i] = []int{0, 0}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				y := min(dp[i-1][j-1][0], dp[i-1][j][0]) + 1
				x := min(dp[i-1][j-1][1], dp[i][j-1][1]) + 1
				dp[i][j] = []int{y, x}

				if x*y > maxArea {
					maxArea = x * y
				}
			} else {
				dp[i][j] = []int{0, 0}
			}

		}
	}

	return maxArea
}

func TestMaximalRectangle(t *testing.T) {
	fmt.Println(maximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	fmt.Println(maximalRectangle([][]byte{}))
	fmt.Println(maximalRectangle([][]byte{{'0'}}))
	fmt.Println(maximalRectangle([][]byte{{'1'}}))

}
