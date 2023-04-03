package hot100

import (
	"fmt"
	"testing"
)

/**
个人想法
先对角线二分查找，再分别行和列二分查找 （错的，这个规律不对，看看用例中的15，不符合这个规律）
从起始行遍历到最大行，从起始列遍历到最大列

虽然是自己写的，但感觉写得太复杂太乱了

时间
28 ms
击败
7.45%
内存
6.6 MB
击败
32.2%


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

	if matrix[startY][startX] == target {
		return true
	}
	if matrix[endY][endX] == target {
		return true
	}

	for endX-startX > 1 || endY-startY > 1 {

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
	x := endX

	for x < len(matrix[0]) {
		startYY, endYY := 0, endY
		for startYY < endYY {

			if endYY-startYY <= 1 {
				if matrix[startYY][x] == target {
					return true
				}
				if matrix[endYY][x] == target {
					return true
				}
				break
			}

			mid := (endYY + startYY) / 2
			if matrix[mid][x] == target {
				return true
			}
			if matrix[mid][x] > target {
				endYY = mid
				continue
			}
			if matrix[mid][x] < target {
				startYY = mid
				continue
			}
		}
		x++
	}

	y := endY
	for y < len(matrix) {
		startXX, endXX := 0, endX
		for startXX < endXX {

			if endXX-startXX <= 1 {
				if matrix[y][startXX] == target {
					return true
				}
				if matrix[y][endXX] == target {
					return true
				}
				break
			}

			mid := (endXX + startXX) / 2
			if matrix[y][mid] == target {
				return true
			}
			if matrix[y][mid] > target {
				endXX = mid
				continue
			}
			if matrix[y][mid] < target {
				startXX = mid
				continue
			}
		}
		y++
	}

	return false
}

func TestSearchMatrix(t *testing.T) {
	//fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5))
	//fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
	//fmt.Println(searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 4))

	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(matrix, 4))
	checkMatrix(matrix, t)

	matrix2 := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}, {10, 13, 14}, {18, 21, 23}}
	fmt.Println(searchMatrix(matrix2, 4))
	checkMatrix(matrix2, t)

	matrix3 := [][]int{{1}, {2}, {3}, {10}, {18}}
	fmt.Println(searchMatrix(matrix3, 4))
	checkMatrix(matrix3, t)

	matrix4 := [][]int{{1}}
	fmt.Println(searchMatrix(matrix4, 4))
	checkMatrix(matrix4, t)

}

func checkMatrix(matrix [][]int, t *testing.T) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if !searchMatrix(matrix, matrix[i][j]) {
				t.Fatal("search error,", matrix[i][j])
			}
		}
	}
}
