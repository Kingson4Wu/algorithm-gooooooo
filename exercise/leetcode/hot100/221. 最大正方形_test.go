package hot100

import (
	"fmt"
	"testing"
)

/**
看过题解做的
dp(i,j)表示以 (i,j) 为右下角，且只包含 1 的正方形的边长最大值

dp(i,j)=min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1

时间
4 ms
击败
96.59%
内存
6.5 MB
击败
13.52%


我们用 dp(i,j)表示以 (i,j) 为右下角，且只包含 1 的正方形的边长最大值。如果我们能计算出所有 dp(i,j)的值，那么其中的最大值即为矩阵中只包含 1 的正方形的边长最大值，其平方即为最大正方形的面积。

那么如何计算 dp 中的每个元素值呢？对于每个位置 (i,j)，检查在矩阵中该位置的值：

如果该位置的值是 0，则 dp(i,j)=0，因为当前位置不可能在由 1 组成的正方形中；

如果该位置的值是 1，则 dp(i,j)的值由其上方、左方和左上方的三个相邻位置的 dp 值决定。具体而言，当前位置的元素值等于三个相邻位置的元素中的最小值加 1，状态转移方程如下：

dp(i,j)=min(dp(i−1, j), dp(i−1, j−1), dp(i, j−1))+1

如果读者对这个状态转移方程感到不解，可以参考 1277. 统计全为 1 的正方形子矩阵的官方题解，其中给出了详细的证明。

此外，还需要考虑边界条件。如果 i 和 j 中至少有一个为 0，则以位置 (i,j) 为右下角的最大正方形的边长只能是 1，因此 dp(i,j)=1。

作者：力扣官方题解
链接：https://leetcode.cn/problems/maximal-square/solutions/234964/zui-da-zheng-fang-xing-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
/*
*
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

输入：matrix = [['0','1'],['1','0']]
输出：1
示例 3：

输入：matrix = [['0']]
输出：0

提示：

m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] 为 '0' 或 '1'
*/
func maximalSquare(matrix [][]byte) int {

	maxLength := 0

	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	/** 初始化 */
	for x := 0; x < len(matrix[0]); x++ {
		if matrix[0][x] == '1' {
			dp[0][x] = 1
		}
		if dp[0][x] > maxLength {
			maxLength = dp[0][x]
		}
	}
	for y := 0; y < len(matrix); y++ {
		if matrix[y][0] == '1' {
			dp[y][0] = 1
		}
		if dp[y][0] > maxLength {
			maxLength = dp[y][0]
		}
	}

	for y := 1; y < len(matrix); y++ {
		for x := 1; x < len(matrix[y]); x++ {
			if matrix[y][x] == '1' {
				dp[y][x] = min(min(dp[y][x-1], dp[y-1][x]), dp[y-1][x-1]) + 1
			} else {
				dp[y][x] = 0
			}
			if dp[y][x] > maxLength {
				maxLength = dp[y][x]
			}
		}
	}
	return maxLength * maxLength
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func TestMaximalSquare(t *testing.T) {
	fmt.Println(maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}
