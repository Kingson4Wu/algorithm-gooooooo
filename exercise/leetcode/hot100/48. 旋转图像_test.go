package hot100

import (
	"fmt"
	"testing"
)

/**
总结规律即可做出来
需要耐心

注意二维数组的第一个下标是第几行，即y，不是x

给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]
输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

时间
0 ms
击败
100%
内存
2.1 MB
击败
34.37%
*/

func rotate(matrix [][]int) {

	if len(matrix) <= 1 {
		return
	}

	for y := 0; y < len(matrix)/2; y++ {
		length := len(matrix) - 2*y
		if length <= 1 {
			break
		}
		x := y
		for i := 0; i < length-1; i++ {
			temp := matrix[y][x+i]
			matrix[y+i][x+length-1], temp = temp, matrix[y+i][x+length-1]
			matrix[y+length-1][x+length-1-i], temp = temp, matrix[y+length-1][x+length-1-i]
			matrix[y+length-1-i][x], temp = temp, matrix[y+length-1-i][x]
			matrix[y][x+i], temp = temp, matrix[y][x+i]
		}

	}

}

func TestRotate(t *testing.T) {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(arr)
	fmt.Println(arr)

	arr2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(arr2)
	fmt.Println(arr2)

	arr3 := [][]int{{5, 1}, {2, 4}}
	rotate(arr3)
	fmt.Println(arr3)

	arr4 := [][]int{{5}}
	rotate(arr4)
	fmt.Println(arr4)
}
