package tencent

import (
	"fmt"
	"testing"
)

/**
自己写的，不过感觉有点复杂

时间
0 ms
击败
100%
内存
2 MB
击败
32.52%
*/

func generateMatrix(n int) [][]int {

	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	num := 1
	length := n - 1

	for round := 1; round <= (n+1)/2; round++ {

		/** top */
		for i := round - 1; i < round-1+length; i++ {
			matrix[round-1][i] = num
			num++
		}

		/** right */
		for i := round - 1; i < round-1+length; i++ {
			matrix[i][n-round] = num
			num++
		}

		/** bottom */
		for i := round - 1; i < round-1+length; i++ {
			matrix[n-round][n-1-i] = num
			num++
		}

		/** left */
		for i := round - 1; i < round-1+length; i++ {
			matrix[n-1-i][round-1] = num
			num++
		}

		length -= 2
	}
	if n%2 != 0 {
		matrix[n/2][n/2] = n * n
	}

	return matrix

}

func TestGenerateMatrix(t *testing.T) {
	matrix := generateMatrix(1)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
