package hot100

import (
	"fmt"
	"testing"
)

/**
个人想法
先对角线二分查找，再分别行和列二分查找
*/

/**
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
输出：false


*/

func searchMatrix(matrix [][]int, target int) bool {

	/** 对角线查找 */
	startX, startY := 0, 0
	endX, endY := len(matrix[0])-1, len(matrix)-1

	if matrix[startX][startY] == target {
		return true
	}
	if matrix[endY][endX] == target {
		return true
	}

	for endX-startX != 1 && endY-startY != 1 {

		midX := (startX + endX) / 2
		midY := (startY + endY) / 2

		if matrix[midY][midX] == target {
			return true
		}
		if matrix[midY][midX] > target {
			endX, endY = midX, midY
			continue
		}
		if matrix[midY][midX] < target {
			startX, startY = midX, midY
			continue
		}
	}

	/** 行和列二分查找 */
	startYY, endYY := 0, endY

	for startYY < endYY {

		if endYY-startYY <= 2 {
			if matrix[startYY][endX] == target {
				return true
			}
			if matrix[endYY][endX] == target {
				return true
			}
			break
		}

		mid := (endYY + startYY) / 2
		if matrix[mid][endX] == target {
			return true
		}
		if matrix[mid][endX] > target {
			endYY = mid
			continue
		}
		if matrix[mid][endX] < target {
			startYY = mid
			continue
		}
	}

	startXX, endXX := 0, endX
	for startXX < endXX {

		if endXX-startXX <= 2 {
			if matrix[endYY][startXX] == target {
				return true
			}
			if matrix[endYY][endXX] == target {
				return true
			}
			break
		}

		mid := (endXX + startXX) / 2
		if matrix[endY][mid] == target {
			return true
		}
		if matrix[endY][mid] > target {
			endXX = mid
			continue
		}
		if matrix[endY][mid] < target {
			startXX = mid
			continue
		}
	}

	return false
}

func TestSearchMatrix(t *testing.T) {
	//fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	//fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
	//fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 4))

	fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 15))

	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if !searchMatrix(matrix, matrix[i][j]) {
				t.Fatal("search error,", matrix[i][j])
			}
		}
	}

}
